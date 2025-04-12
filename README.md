# MarkDown API
 - This is a Perticular API consist of 4 endpoint : 
	 - /ADD: To add a file
     -  /GetFile :  TO get files which are already added  
     - /GetAllFiles :  To List all the files   
     - /GetRenderedFiles - To recieve Files which are converted to HTML.
 
-  Here you will have to add a module of Name Helper , 
- This module will host all the meta data related to api in it.
	- This Module will host the following metadata:
		- FileHeaderName: Represent a name of the file tag in the form data.
		- FileStoreLocation : Represent the path where files will stored. 
	 
- Also we will need to create a folder of the name stated in the meta-Data module, to store the files.

Pending Task : 
- Instead of Metadata module there should be environment variables
- files to be store in a db.
