package model

type ReleaseValues struct {
	ClusterDomain string `yaml:"clusterDomain"`
	Config        struct {
		ServerConfigStrictValidationEnabled string `yaml:"server.config.strict_validation.enabled"`
	} `yaml:"config"`
	Env struct {
	} `yaml:"env"`
	Image struct {
		CustomImage     string `yaml:"customImage"`
		ImagePullPolicy string `yaml:"imagePullPolicy"`
	} `yaml:"image"`
	Jvm struct {
		AdditionalJvmArguments      []interface{} `yaml:"additionalJvmArguments"`
		UseNeo4JDefaultJvmArguments bool          `yaml:"useNeo4jDefaultJvmArguments"`
	} `yaml:"jvm"`
	LivenessProbe struct {
		FailureThreshold int `yaml:"failureThreshold"`
		PeriodSeconds    int `yaml:"periodSeconds"`
		TimeoutSeconds   int `yaml:"timeoutSeconds"`
	} `yaml:"livenessProbe"`
	LogInitialPassword bool `yaml:"logInitialPassword"`
	Logs               struct {
		ServerLogsXML string `yaml:"server_logs_xml"`
		UserLogsXML   string `yaml:"user_logs_xml"`
	} `yaml:"logs"`
	Neo4J struct {
		AcceptLicenseAgreement        string      `yaml:"acceptLicenseAgreement"`
		Edition                       string      `yaml:"edition"`
		Labels                        interface{} `yaml:"labels"`
		Name                          string      `yaml:"name"`
		OfflineMaintenanceModeEnabled bool        `yaml:"offlineMaintenanceModeEnabled"`
		Password                      string      `yaml:"password"`
		Resources                     struct {
			CPU    string `yaml:"cpu"`
			Limits struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"limits"`
			Memory   string `yaml:"memory"`
			Requests struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"requests"`
		} `yaml:"resources"`
	} `yaml:"neo4j"`
	NodeSelector interface{} `yaml:"nodeSelector"`
	PodSpec      struct {
		Annotations                   interface{}   `yaml:"annotations"`
		Containers                    []interface{} `yaml:"containers"`
		InitContainers                []interface{} `yaml:"initContainers"`
		Loadbalancer                  string        `yaml:"loadbalancer"`
		NodeAffinity                  interface{}   `yaml:"nodeAffinity"`
		PodAntiAffinity               bool          `yaml:"podAntiAffinity"`
		PriorityClassName             string        `yaml:"priorityClassName"`
		ServiceAccountName            string        `yaml:"serviceAccountName"`
		TerminationGracePeriodSeconds int           `yaml:"terminationGracePeriodSeconds"`
		Tolerations                   interface{}   `yaml:"tolerations"`
	} `yaml:"podSpec"`
	ReadinessProbe struct {
		FailureThreshold int `yaml:"failureThreshold"`
		PeriodSeconds    int `yaml:"periodSeconds"`
		TimeoutSeconds   int `yaml:"timeoutSeconds"`
	} `yaml:"readinessProbe"`
	SecurityContext struct {
		FsGroup             int    `yaml:"fsGroup"`
		FsGroupChangePolicy string `yaml:"fsGroupChangePolicy"`
		RunAsGroup          int    `yaml:"runAsGroup"`
		RunAsNonRoot        bool   `yaml:"runAsNonRoot"`
		RunAsUser           int    `yaml:"runAsUser"`
	} `yaml:"securityContext"`
	Services struct {
		Admin struct {
			Annotations struct {
			} `yaml:"annotations"`
			Enabled bool `yaml:"enabled"`
			Spec    struct {
				Type string `yaml:"type"`
			} `yaml:"spec"`
		} `yaml:"admin"`
		Default struct {
			Annotations struct {
			} `yaml:"annotations"`
		} `yaml:"default"`
		Internals struct {
			Annotations struct {
			} `yaml:"annotations"`
			Enabled bool `yaml:"enabled"`
		} `yaml:"internals"`
		Neo4J struct {
			Annotations struct {
			} `yaml:"annotations"`
			Enabled      bool `yaml:"enabled"`
			MultiCluster bool `yaml:"multiCluster"`
			Ports        struct {
				Backup struct {
					Enabled bool `yaml:"enabled"`
				} `yaml:"backup"`
				Bolt struct {
					Enabled bool `yaml:"enabled"`
				} `yaml:"bolt"`
				HTTP struct {
					Enabled bool `yaml:"enabled"`
				} `yaml:"http"`
				HTTPS struct {
					Enabled bool `yaml:"enabled"`
				} `yaml:"https"`
			} `yaml:"ports"`
			Selector struct {
				HelmNeo4JComNeo4JLoadbalancer string `yaml:"helm.neo4j.com/neo4j.loadbalancer"`
			} `yaml:"selector"`
			Spec struct {
				Type string `yaml:"type"`
			} `yaml:"spec"`
		} `yaml:"neo4j"`
	} `yaml:"services"`
	Ssl struct {
		Bolt struct {
			PrivateKey struct {
				SecretName interface{} `yaml:"secretName"`
				SubPath    interface{} `yaml:"subPath"`
			} `yaml:"privateKey"`
			PublicCertificate struct {
				SecretName interface{} `yaml:"secretName"`
				SubPath    interface{} `yaml:"subPath"`
			} `yaml:"publicCertificate"`
			RevokedCerts struct {
				Sources []interface{} `yaml:"sources"`
			} `yaml:"revokedCerts"`
			TrustedCerts struct {
				Sources []interface{} `yaml:"sources"`
			} `yaml:"trustedCerts"`
		} `yaml:"bolt"`
		HTTPS struct {
			PrivateKey struct {
				SecretName interface{} `yaml:"secretName"`
				SubPath    interface{} `yaml:"subPath"`
			} `yaml:"privateKey"`
			PublicCertificate struct {
				SecretName interface{} `yaml:"secretName"`
				SubPath    interface{} `yaml:"subPath"`
			} `yaml:"publicCertificate"`
			RevokedCerts struct {
				Sources []interface{} `yaml:"sources"`
			} `yaml:"revokedCerts"`
			TrustedCerts struct {
				Sources []interface{} `yaml:"sources"`
			} `yaml:"trustedCerts"`
		} `yaml:"https"`
	} `yaml:"ssl"`
	StartupProbe struct {
		FailureThreshold int `yaml:"failureThreshold"`
		PeriodSeconds    int `yaml:"periodSeconds"`
	} `yaml:"startupProbe"`
	Statefulset struct {
		Metadata struct {
			Annotations interface{} `yaml:"annotations"`
		} `yaml:"metadata"`
	} `yaml:"statefulset"`
	Volumes struct {
		Backups struct {
			Mode  string `yaml:"mode"`
			Share struct {
				Name string `yaml:"name"`
			} `yaml:"share"`
		} `yaml:"backups"`
		Data struct {
			DefaultStorageClass struct {
				AccessModes []string `yaml:"accessModes"`
				Requests    struct {
					Storage string `yaml:"storage"`
				} `yaml:"requests"`
			} `yaml:"defaultStorageClass"`
			Dynamic struct {
				AccessModes []string `yaml:"accessModes"`
				Requests    struct {
					Storage string `yaml:"storage"`
				} `yaml:"requests"`
				StorageClassName string `yaml:"storageClassName"`
			} `yaml:"dynamic"`
			Mode     string `yaml:"mode"`
			Selector struct {
				AccessModes []string `yaml:"accessModes"`
				Requests    struct {
					Storage string `yaml:"storage"`
				} `yaml:"requests"`
				SelectorTemplate struct {
					MatchLabels struct {
						App                    string `yaml:"app"`
						HelmNeo4JComVolumeRole string `yaml:"helm.neo4j.com/volume-role"`
					} `yaml:"matchLabels"`
				} `yaml:"selectorTemplate"`
				StorageClassName string `yaml:"storageClassName"`
			} `yaml:"selector"`
			Volume struct {
				SetOwnerAndGroupWritableFilePermissions bool `yaml:"setOwnerAndGroupWritableFilePermissions"`
			} `yaml:"volume"`
			VolumeClaimTemplate struct {
			} `yaml:"volumeClaimTemplate"`
		} `yaml:"data"`
		Import struct {
			Mode  string `yaml:"mode"`
			Share struct {
				Name string `yaml:"name"`
			} `yaml:"share"`
		} `yaml:"import"`
		Licenses struct {
			Mode  string `yaml:"mode"`
			Share struct {
				Name string `yaml:"name"`
			} `yaml:"share"`
		} `yaml:"licenses"`
		Logs struct {
			Mode  string `yaml:"mode"`
			Share struct {
				Name string `yaml:"name"`
			} `yaml:"share"`
		} `yaml:"logs"`
		Metrics struct {
			Mode  string `yaml:"mode"`
			Share struct {
				Name string `yaml:"name"`
			} `yaml:"share"`
		} `yaml:"metrics"`
	} `yaml:"volumes"`
}
