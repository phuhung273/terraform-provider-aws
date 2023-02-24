// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/service/accessanalyzer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/account"
	"github.com/hashicorp/terraform-provider-aws/internal/service/acm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/acmpca"
	"github.com/hashicorp/terraform-provider-aws/internal/service/amp"
	"github.com/hashicorp/terraform-provider-aws/internal/service/amplify"
	"github.com/hashicorp/terraform-provider-aws/internal/service/apigateway"
	"github.com/hashicorp/terraform-provider-aws/internal/service/apigatewayv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appautoscaling"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appconfig"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appflow"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appintegrations"
	"github.com/hashicorp/terraform-provider-aws/internal/service/applicationinsights"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appmesh"
	"github.com/hashicorp/terraform-provider-aws/internal/service/apprunner"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appstream"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appsync"
	"github.com/hashicorp/terraform-provider-aws/internal/service/athena"
	"github.com/hashicorp/terraform-provider-aws/internal/service/auditmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/autoscaling"
	"github.com/hashicorp/terraform-provider-aws/internal/service/autoscalingplans"
	"github.com/hashicorp/terraform-provider-aws/internal/service/backup"
	"github.com/hashicorp/terraform-provider-aws/internal/service/batch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/budgets"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ce"
	"github.com/hashicorp/terraform-provider-aws/internal/service/chime"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloud9"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudcontrol"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudformation"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudfront"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudhsmv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudsearch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudtrail"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudwatch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codeartifact"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codebuild"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codecommit"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codepipeline"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codestarconnections"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codestarnotifications"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cognitoidentity"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cognitoidp"
	"github.com/hashicorp/terraform-provider-aws/internal/service/comprehend"
	"github.com/hashicorp/terraform-provider-aws/internal/service/computeoptimizer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/configservice"
	"github.com/hashicorp/terraform-provider-aws/internal/service/connect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/controltower"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cur"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dataexchange"
	"github.com/hashicorp/terraform-provider-aws/internal/service/datapipeline"
	"github.com/hashicorp/terraform-provider-aws/internal/service/datasync"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dax"
	"github.com/hashicorp/terraform-provider-aws/internal/service/deploy"
	"github.com/hashicorp/terraform-provider-aws/internal/service/detective"
	"github.com/hashicorp/terraform-provider-aws/internal/service/devicefarm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/directconnect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dlm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dms"
	"github.com/hashicorp/terraform-provider-aws/internal/service/docdb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ds"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dynamodb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ecr"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ecrpublic"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ecs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/efs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/eks"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elasticache"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elasticbeanstalk"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elasticsearch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elastictranscoder"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elbv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/emr"
	"github.com/hashicorp/terraform-provider-aws/internal/service/emrcontainers"
	"github.com/hashicorp/terraform-provider-aws/internal/service/emrserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/service/events"
	"github.com/hashicorp/terraform-provider-aws/internal/service/evidently"
	"github.com/hashicorp/terraform-provider-aws/internal/service/firehose"
	"github.com/hashicorp/terraform-provider-aws/internal/service/fis"
	"github.com/hashicorp/terraform-provider-aws/internal/service/fms"
	"github.com/hashicorp/terraform-provider-aws/internal/service/fsx"
	"github.com/hashicorp/terraform-provider-aws/internal/service/gamelift"
	"github.com/hashicorp/terraform-provider-aws/internal/service/glacier"
	"github.com/hashicorp/terraform-provider-aws/internal/service/globalaccelerator"
	"github.com/hashicorp/terraform-provider-aws/internal/service/glue"
	"github.com/hashicorp/terraform-provider-aws/internal/service/grafana"
	"github.com/hashicorp/terraform-provider-aws/internal/service/greengrass"
	"github.com/hashicorp/terraform-provider-aws/internal/service/guardduty"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iam"
	"github.com/hashicorp/terraform-provider-aws/internal/service/identitystore"
	"github.com/hashicorp/terraform-provider-aws/internal/service/imagebuilder"
	"github.com/hashicorp/terraform-provider-aws/internal/service/inspector"
	"github.com/hashicorp/terraform-provider-aws/internal/service/inspector2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iot"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iotanalytics"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iotevents"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ivs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ivschat"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kafka"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kafkaconnect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kendra"
	"github.com/hashicorp/terraform-provider-aws/internal/service/keyspaces"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesis"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesisanalytics"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesisanalyticsv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesisvideo"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kms"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lakeformation"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lambda"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lexmodels"
	"github.com/hashicorp/terraform-provider-aws/internal/service/licensemanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lightsail"
	"github.com/hashicorp/terraform-provider-aws/internal/service/location"
	"github.com/hashicorp/terraform-provider-aws/internal/service/logs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/macie"
	"github.com/hashicorp/terraform-provider-aws/internal/service/macie2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediaconnect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediaconvert"
	"github.com/hashicorp/terraform-provider-aws/internal/service/medialive"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediapackage"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediastore"
	"github.com/hashicorp/terraform-provider-aws/internal/service/memorydb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/meta"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mq"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mwaa"
	"github.com/hashicorp/terraform-provider-aws/internal/service/neptune"
	"github.com/hashicorp/terraform-provider-aws/internal/service/networkfirewall"
	"github.com/hashicorp/terraform-provider-aws/internal/service/networkmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/oam"
	"github.com/hashicorp/terraform-provider-aws/internal/service/opensearch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/opensearchserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/service/opsworks"
	"github.com/hashicorp/terraform-provider-aws/internal/service/organizations"
	"github.com/hashicorp/terraform-provider-aws/internal/service/outposts"
	"github.com/hashicorp/terraform-provider-aws/internal/service/pinpoint"
	"github.com/hashicorp/terraform-provider-aws/internal/service/pipes"
	"github.com/hashicorp/terraform-provider-aws/internal/service/pricing"
	"github.com/hashicorp/terraform-provider-aws/internal/service/qldb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/quicksight"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ram"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rds"
	"github.com/hashicorp/terraform-provider-aws/internal/service/redshift"
	"github.com/hashicorp/terraform-provider-aws/internal/service/redshiftdata"
	"github.com/hashicorp/terraform-provider-aws/internal/service/redshiftserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/service/resourceexplorer2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/resourcegroups"
	"github.com/hashicorp/terraform-provider-aws/internal/service/resourcegroupstaggingapi"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rolesanywhere"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53domains"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53recoverycontrolconfig"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53recoveryreadiness"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53resolver"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rum"
	"github.com/hashicorp/terraform-provider-aws/internal/service/s3"
	"github.com/hashicorp/terraform-provider-aws/internal/service/s3control"
	"github.com/hashicorp/terraform-provider-aws/internal/service/s3outposts"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sagemaker"
	"github.com/hashicorp/terraform-provider-aws/internal/service/scheduler"
	"github.com/hashicorp/terraform-provider-aws/internal/service/schemas"
	"github.com/hashicorp/terraform-provider-aws/internal/service/secretsmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/securityhub"
	"github.com/hashicorp/terraform-provider-aws/internal/service/serverlessrepo"
	"github.com/hashicorp/terraform-provider-aws/internal/service/servicecatalog"
	"github.com/hashicorp/terraform-provider-aws/internal/service/servicediscovery"
	"github.com/hashicorp/terraform-provider-aws/internal/service/servicequotas"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ses"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sesv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sfn"
	"github.com/hashicorp/terraform-provider-aws/internal/service/shield"
	"github.com/hashicorp/terraform-provider-aws/internal/service/signer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/simpledb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sns"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sqs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssoadmin"
	"github.com/hashicorp/terraform-provider-aws/internal/service/storagegateway"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sts"
	"github.com/hashicorp/terraform-provider-aws/internal/service/swf"
	"github.com/hashicorp/terraform-provider-aws/internal/service/synthetics"
	"github.com/hashicorp/terraform-provider-aws/internal/service/timestreamwrite"
	"github.com/hashicorp/terraform-provider-aws/internal/service/transcribe"
	"github.com/hashicorp/terraform-provider-aws/internal/service/transfer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/waf"
	"github.com/hashicorp/terraform-provider-aws/internal/service/wafregional"
	"github.com/hashicorp/terraform-provider-aws/internal/service/wafv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/worklink"
	"github.com/hashicorp/terraform-provider-aws/internal/service/workspaces"
	"github.com/hashicorp/terraform-provider-aws/internal/service/xray"
	"golang.org/x/exp/slices"
)

func servicePackages(context.Context) []conns.ServicePackage {
	v := []conns.ServicePackage{
		accessanalyzer.ServicePackage,
		account.ServicePackage,
		acm.ServicePackage,
		acmpca.ServicePackage,
		amp.ServicePackage,
		amplify.ServicePackage,
		apigateway.ServicePackage,
		apigatewayv2.ServicePackage,
		appautoscaling.ServicePackage,
		appconfig.ServicePackage,
		appflow.ServicePackage,
		appintegrations.ServicePackage,
		applicationinsights.ServicePackage,
		appmesh.ServicePackage,
		apprunner.ServicePackage,
		appstream.ServicePackage,
		appsync.ServicePackage,
		athena.ServicePackage,
		auditmanager.ServicePackage,
		autoscaling.ServicePackage,
		autoscalingplans.ServicePackage,
		backup.ServicePackage,
		batch.ServicePackage,
		budgets.ServicePackage,
		ce.ServicePackage,
		chime.ServicePackage,
		cloud9.ServicePackage,
		cloudcontrol.ServicePackage,
		cloudformation.ServicePackage,
		cloudfront.ServicePackage,
		cloudhsmv2.ServicePackage,
		cloudsearch.ServicePackage,
		cloudtrail.ServicePackage,
		cloudwatch.ServicePackage,
		codeartifact.ServicePackage,
		codebuild.ServicePackage,
		codecommit.ServicePackage,
		codepipeline.ServicePackage,
		codestarconnections.ServicePackage,
		codestarnotifications.ServicePackage,
		cognitoidentity.ServicePackage,
		cognitoidp.ServicePackage,
		comprehend.ServicePackage,
		computeoptimizer.ServicePackage,
		configservice.ServicePackage,
		connect.ServicePackage,
		controltower.ServicePackage,
		cur.ServicePackage,
		dataexchange.ServicePackage,
		datapipeline.ServicePackage,
		datasync.ServicePackage,
		dax.ServicePackage,
		deploy.ServicePackage,
		detective.ServicePackage,
		devicefarm.ServicePackage,
		directconnect.ServicePackage,
		dlm.ServicePackage,
		dms.ServicePackage,
		docdb.ServicePackage,
		ds.ServicePackage,
		dynamodb.ServicePackage,
		ec2.ServicePackage,
		ecr.ServicePackage,
		ecrpublic.ServicePackage,
		ecs.ServicePackage,
		efs.ServicePackage,
		eks.ServicePackage,
		elasticache.ServicePackage,
		elasticbeanstalk.ServicePackage,
		elasticsearch.ServicePackage,
		elastictranscoder.ServicePackage,
		elb.ServicePackage,
		elbv2.ServicePackage,
		emr.ServicePackage,
		emrcontainers.ServicePackage,
		emrserverless.ServicePackage,
		events.ServicePackage,
		evidently.ServicePackage,
		firehose.ServicePackage,
		fis.ServicePackage,
		fms.ServicePackage,
		fsx.ServicePackage,
		gamelift.ServicePackage,
		glacier.ServicePackage,
		globalaccelerator.ServicePackage,
		glue.ServicePackage,
		grafana.ServicePackage,
		greengrass.ServicePackage,
		guardduty.ServicePackage,
		iam.ServicePackage,
		identitystore.ServicePackage,
		imagebuilder.ServicePackage,
		inspector.ServicePackage,
		inspector2.ServicePackage,
		iot.ServicePackage,
		iotanalytics.ServicePackage,
		iotevents.ServicePackage,
		ivs.ServicePackage,
		ivschat.ServicePackage,
		kafka.ServicePackage,
		kafkaconnect.ServicePackage,
		kendra.ServicePackage,
		keyspaces.ServicePackage,
		kinesis.ServicePackage,
		kinesisanalytics.ServicePackage,
		kinesisanalyticsv2.ServicePackage,
		kinesisvideo.ServicePackage,
		kms.ServicePackage,
		lakeformation.ServicePackage,
		lambda.ServicePackage,
		lexmodels.ServicePackage,
		licensemanager.ServicePackage,
		lightsail.ServicePackage,
		location.ServicePackage,
		logs.ServicePackage,
		macie.ServicePackage,
		macie2.ServicePackage,
		mediaconnect.ServicePackage,
		mediaconvert.ServicePackage,
		medialive.ServicePackage,
		mediapackage.ServicePackage,
		mediastore.ServicePackage,
		memorydb.ServicePackage,
		meta.ServicePackage,
		mq.ServicePackage,
		mwaa.ServicePackage,
		neptune.ServicePackage,
		networkfirewall.ServicePackage,
		networkmanager.ServicePackage,
		oam.ServicePackage,
		opensearch.ServicePackage,
		opensearchserverless.ServicePackage,
		opsworks.ServicePackage,
		organizations.ServicePackage,
		outposts.ServicePackage,
		pinpoint.ServicePackage,
		pipes.ServicePackage,
		pricing.ServicePackage,
		qldb.ServicePackage,
		quicksight.ServicePackage,
		ram.ServicePackage,
		rds.ServicePackage,
		redshift.ServicePackage,
		redshiftdata.ServicePackage,
		redshiftserverless.ServicePackage,
		resourceexplorer2.ServicePackage,
		resourcegroups.ServicePackage,
		resourcegroupstaggingapi.ServicePackage,
		rolesanywhere.ServicePackage,
		route53.ServicePackage,
		route53domains.ServicePackage,
		route53recoverycontrolconfig.ServicePackage,
		route53recoveryreadiness.ServicePackage,
		route53resolver.ServicePackage,
		rum.ServicePackage,
		s3.ServicePackage,
		s3control.ServicePackage,
		s3outposts.ServicePackage,
		sagemaker.ServicePackage,
		scheduler.ServicePackage,
		schemas.ServicePackage,
		secretsmanager.ServicePackage,
		securityhub.ServicePackage,
		serverlessrepo.ServicePackage,
		servicecatalog.ServicePackage,
		servicediscovery.ServicePackage,
		servicequotas.ServicePackage,
		ses.ServicePackage,
		sesv2.ServicePackage,
		sfn.ServicePackage,
		shield.ServicePackage,
		signer.ServicePackage,
		simpledb.ServicePackage,
		sns.ServicePackage,
		sqs.ServicePackage,
		ssm.ServicePackage,
		ssoadmin.ServicePackage,
		storagegateway.ServicePackage,
		sts.ServicePackage,
		swf.ServicePackage,
		synthetics.ServicePackage,
		timestreamwrite.ServicePackage,
		transcribe.ServicePackage,
		transfer.ServicePackage,
		waf.ServicePackage,
		wafregional.ServicePackage,
		wafv2.ServicePackage,
		worklink.ServicePackage,
		workspaces.ServicePackage,
		xray.ServicePackage,
	}

	return slices.Clone(v)
}
