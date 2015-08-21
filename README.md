# Simple Golang Scaling Demo on PWS / Cloud Foundry

This is a simple Golang application deployed on Pivotal Web Services. The application is written using the martini framework. It is best viewed side by side in the apps manager web console to see the ease of scaling and recovery. Try the following:

* scale an application using cf scale from the command line or scale using the apps manager webconsole. (http://console.run.pivotal.io)
* Kill an instance using the kill link below. See how the events are captured on the console and how it automatically recovers.
* Return to the app root in the browser and see that you can still access the application even though an instance has crashed
* Look at the stats about the current instance below. Pivotal Cloud Foundry deploys the application across multiple VMs automatically.

Note: Don't have a PWS account, get a free 60 day trial at http://run.pivotal.io/register


