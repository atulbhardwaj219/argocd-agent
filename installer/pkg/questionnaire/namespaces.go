package questionnaire

import (
	"github.com/codefresh-io/argocd-listener/installer/pkg/install"
	"github.com/codefresh-io/argocd-listener/installer/pkg/kube"
	"github.com/codefresh-io/argocd-listener/installer/pkg/prompt"
)

func AskAboutNamespace(kubeOptions *install.Kube, kubeClient kube.Kube, setDefaultNamespace bool) error {
	const defaultNamespace = "argocd"
	if kubeOptions.Namespace != "" {
		return nil
	}

	namespaces, err := kubeClient.GetNamespaces()
	if err != nil {
		err = prompt.InputWithDefault(&kubeOptions.Namespace, "Kubernetes namespace to update", "default")
		if err != nil {
			return err
		}
	} else {
		if setDefaultNamespace {
			for _, namespace := range namespaces {
				if namespace == defaultNamespace {
					kubeOptions.Namespace = defaultNamespace
					return nil
				}
			}
		}

		err, selectedNamespace := prompt.Select(namespaces, "Select Kubernetes namespace")
		if err != nil {
			return err
		}
		kubeOptions.Namespace = selectedNamespace
	}
	return nil
}