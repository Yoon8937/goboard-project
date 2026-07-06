package post

import (
	// Google Cloud
	_ "cloud.google.com/go/bigquery"
	_ "cloud.google.com/go/pubsub"
	_ "cloud.google.com/go/spanner"
	_ "cloud.google.com/go/storage"

	// Azure
	_ "github.com/Azure/azure-sdk-for-go/sdk/azcore"
	_ "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	_ "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	_ "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

	// AWS
	_ "github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	_ "github.com/aws/aws-sdk-go-v2/service/ec2"
	_ "github.com/aws/aws-sdk-go-v2/service/lambda"
	_ "github.com/aws/aws-sdk-go-v2/service/rds"
	_ "github.com/aws/aws-sdk-go-v2/service/s3"

	// Core & Others
	_ "github.com/cilium/cilium/pkg/api" // 패키지 경로 주의
	_ "github.com/containerd/containerd"
	_ "github.com/crossplane/crossplane-runtime/pkg/logging"
	_ "github.com/docker/docker/client"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	_ "github.com/gin-gonic/gin"
	_ "github.com/hashicorp/consul/api"
	_ "github.com/hashicorp/nomad/api"
	_ "github.com/hashicorp/terraform-exec/tfexec"
	_ "github.com/hashicorp/terraform-plugin-framework/provider"
	_ "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	_ "github.com/hashicorp/vault/api"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/moby/moby/client"
	_ "github.com/prometheus/alertmanager/api"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/prometheus/promql"
	_ "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	// Observability
	_ "go.opentelemetry.io/otel"
	_ "go.opentelemetry.io/otel/sdk/trace"
	_ "go.uber.org/zap"

	// Google & Ecosystem
	_ "google.golang.org/api/option"
	_ "google.golang.org/grpc"
	_ "google.golang.org/protobuf/proto"
	_ "gopkg.in/natefinch/lumberjack.v2"
	_ "gopkg.in/yaml.v3"

	// CNCF & K8s
	_ "helm.sh/helm/v3/pkg/action"
	_ "istio.io/api/networking/v1alpha3"
	_ "istio.io/client-go/pkg/clientset/versioned"
	_ "k8s.io/api/core/v1"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apiserver/pkg/server"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/component-base/version"
	_ "k8s.io/kube-openapi/pkg/common"
	_ "k8s.io/kubectl/pkg/cmd"
	_ "modernc.org/sqlite"
	_ "oras.land/oras-go/v2"
	_ "sigs.k8s.io/controller-runtime"
)

import (
	_ "github.com/gofiber/fiber/v3"
	_ "entgo.io/ent"
)
