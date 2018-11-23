package ssh

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"github.com/shenshouer/k8s-ssh-pod/config"
	"github.com/shenshouer/k8s-ssh-pod/k8s"
	"github.com/shenshouer/k8s-ssh-pod/log"
	"golang.org/x/crypto/ssh/terminal"
)

func StartSSH(conf *config.Config) error {
	ssh.Handle(sshHandler)

	options := []ssh.Option{
		ssh.HostKeyFile(conf.PrivateKey),
		ssh.PasswordAuth(passwordHandler),
	}

	log.Infof("Start SSH Server at %s", conf.ServeAddr)
	return ssh.ListenAndServe(conf.ServeAddr, nil, options...)
}

// passwordHandler Authentication SSH username and password
// TODO: support LDAP
func passwordHandler(ctx ssh.Context, password string) bool {
	log.Infof("==>> SSH User:%s Password: %s", ctx.User(), password)
	return true
}

func sshHandler(s ssh.Session) {
	defer s.Exit(0)
	t := terminal.NewTerminal(s, "")
	for {
		// select namespace
		ns, exit, err := promptAndSelectNames(s, t, listNamespaces)
		if err != nil {
			log.Error(err)
			_, _ = fmt.Fprintf(s, "[ERROR]: %v", err)
			return
		}
		if exit {
			return
		}

		// select pod
		var getPod = func() ([]string, error) { return listPodWithNamespace(ns) }
		selectPod, exit, err := promptAndSelectNames(s, t, getPod)
		if err != nil {
			log.Error(err)
			_, _ = fmt.Fprintf(s, "[ERROR]: %v", err)
			return
		}
		if exit {
			return
		}

		clientset, config, err := k8s.NewK8SClient()
		if err != nil {
			_, _ = fmt.Fprintf(s, "生成集群Config错误:%v\r\n", err)
			return
		}

		// exec
		tsession := k8s.TerminalSession{SSHSession: s}
		if err = k8s.StartProcess(clientset, config, []string{"/bin/sh"}, tsession, ns, selectPod, ""); err != nil {
			log.Error("==>>k8s.StartProcess", err)
			_, _ = fmt.Fprintf(s, "<<==>>k8s.StartProcess:%v\r\n", err)
			s.Exit(0)
			return
		}
		// time.Sleep(10 * time.Second)
	}
}

func listNamespaces() ([]string, error) {
	client, _, err := k8s.NewK8SClient()
	if err != nil {
		return nil, err
	}

	nsList, err := k8s.GetNamespaces(client)
	if err != nil {
		return nil, err
	}
	nsNames := []string{}
	for _, item := range nsList.Items {
		nsNames = append(nsNames, item.Name)
	}
	return nsNames, nil
}

func listPodWithNamespace(ns string) ([]string, error) {
	client, _, err := k8s.NewK8SClient()
	if err != nil {
		return nil, err
	}

	podList, err := k8s.GetPods(client, ns)
	if err != nil {
		return nil, err
	}

	podNames := []string{}
	for _, v := range podList.Items {
		podNames = append(podNames, v.Name)
	}
	return podNames, nil
}
