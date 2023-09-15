import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { RestApi, LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';

export class InputManagementStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here

    // example resource
    // const queue = new sqs.Queue(this, 'InputManagementQueue', {
    //   visibilityTimeout: cdk.Duration.seconds(300)
    // });

    const fn = new lambda.Function(this, "MyLambda", {
      code: lambda.Code.fromAsset('lambdas'),
      handler: "main",
      runtime: lambda.Runtime.PROVIDED_AL2
    })

    const gateway = new RestApi(this, "myGateway", {
      defaultCorsPreflightOptions: {
        allowOrigins: ["*"],
        allowMethods: ['GET', 'POST', 'OPTIONS', 'DELETE', 'PUT'],
      }
    })

    const integration = new LambdaIntegration(fn);
    const testResource = gateway.root.addResource("test");

    testResource.addMethod("GET", integration);
    testResource.addMethod("POST", integration);
  }
}
