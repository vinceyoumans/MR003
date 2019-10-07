# MUELLAR REPORT -System 2 

Digested files.

This is the secondary system to digest and present the MR.

This project includes the following:
Web portal with full text of MR.  Features"
1. Heat map of specific words with clickable blocks to drill down to individual pages.  For example, if you wanted to "light up" all of the pages with the word "election", then visitor just needs to select that word and a map will appear.  Click the hotsopt to dril down to specific pages.

2. SPeach version of MR.  A podcast of the MR.  Visitor selected words will be include a bell sound.

3. Mobile Version that displays large text for reading on mobile phone.


## System one 
converted the original Report, which was 450 pages of images in PDF form, into JSON text format.  This system is done in another Repo.

## System two
The system will index and digest the output into something that can be rendered in other formats such as Audio, Mobile Reader and Heat map.
Analysis is on indivudual words.

## Vision.
It is dificult for most people to get their hands on this document and actually study it in its present format.  My vision is a system that will engage readers using new strategies to communicate and deliver a convincing story.

### Question:
Q: This seems like a lot of work...  Why not use a Lucene / Solr approach?  
A: Lucene/Solr, tomcat, JVM's, Websphere, .NET...  al require heavy hosting platforms that are difficult to scale and have high administration obligations.
My goal is a system that can handle Huge traffic spikes, Fast responce rates to visitors, and Low administration costs.  Based on past expereince, I also expect serious Electronic attacks sponsored by professionals with a paid agenda.  Because of daily exploits, Traditional hosting strategies are simply no longer secure.   


Q: Hosting Sollutions?
My solutions will include:
1. GoLang/Docker Scratch minimalist solutions that scales by adding hosting resoruces.  Using Gorilla MUX.  DB eiter Cloud: Fauna, DigitalOcean REDIS, or one Document store of Azure, AWS or Google.  I may use BOLTdb which I would host in a Docker container.

2. Static WebSite.  GoLang will generate the pages.

3. AWS Lambda behind ALB ( rigged to defend against DOD type attacks )


