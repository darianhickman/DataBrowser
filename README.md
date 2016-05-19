# DataBrowser
Add features to Google Chrome any way possible to browse big data whereever  in whatever format. Excel, SQL, CSV, large JSON, large nested XML to start.

#Eventual platforms.  
Completely hosted in cloud. 
Installed as Desktop app probably using Electron. 
Installed as Intranet server app to share internally and keep app maintenance simple for Enterprise customers.

#Potential Features Tech stack

 For csv browsing PapaParse. 

If features need a web server likely Caddy for simple lightweight or if we need a full featured web app server then Beego.  Both have a lot of adoption. 

If we need a way to visualize how data flows between components JSJoint.  

JSJoint for visualizing/graphical editing of flow. 
http://www.jointjs.com/

Google Dataflow for doing ETL in streaming and batch. 
https://cloud.google.com/dataflow/

Some kind of metadata management standard that exists.  

If we need to offer a message queue solution: 


Moving data between different source:
https://github.com/golang/protobuf

Authenticating to Kerberos:
https://github.com/apcera/gssapi

Message queue support:  
https://github.com/nsqio/nsq

