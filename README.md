# kafka-php-go-distribution

MINI-PROJECT
Using Windows 10 but I'm sure if someone is good with Linux, they could change a few things and make it work too.

1. You'll want to install Wamp to host a local server(basically to run PHP website)
http://www.wampserver.com/en/

2. You'll also want JDK
http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html

3. And you'll want Go Lang just in case but skip for now if you want.
https://golang.org/dl/

4. Once you have Wamp installed and hopefully running (if not, check the issue with port 80 being taken), 
you'll want to download this repository and put in inside your Wamp folder under www like the picture below:
![alt text](https://i.imgur.com/k8vGMNO.png)

5. Run Wamp if you haven't done so and open any browser and go to 'localhost' or 'localhost:port',
hopefully, you'll see the "This is working!"
![alt text](https://i.imgur.com/43Weas1.png)

6. If you're following without problems, head inside the 'kafka_2.12-1.1.1' folder, inside you'll see a COMMANDS.txt
run each command inside a new CMD prompt from the same location as the COMMANDs.txt 
(wait a bit between each... especially for creating the topic)
(for windows, you can shift-right click or type cmd in the fold url to open a new prompt, otherwise use cd *folder route*)
![alt text](https://i.imgur.com/OAZRF6U.png)

7. When you are done running all of those, you should have something like this:
(don't close any of them)
(I already had the topic created, but you'll get yours created)
![alt text](https://i.imgur.com/QVki1Uz.png)

8. Next go back to the www folder but this time head into the GoListener folder
(open another CMD prompt but this time at the GoListener folder, run 'GoListener.exe')
Let it run in the background (don't close any cmd prompt opened so far)
![alt text](https://i.imgur.com/ClHYbnR.png)

9. using whatever software/code you can find to POST JSON, I'll be using Postman(extension for chrome)
select a post command and the URL 'localhost/ingest.php' with the same JSON shown below
(if using Postman, select the 'Body' tab, choose 'raw' and drop down to JSON, like the picture)
![alt text](https://i.imgur.com/3TssjK4.png)

10. Hit Send or however you'll send JSON over. In a little bit, you should see incoming messages sent from Kafka over to GoListener.
(Response has sample JSON given plus a personal Get-Request to one of my APIs)
(if you don't see any message try restarting GoListener.exe in cmd or waiting just a little bit longer)
![alt text](https://i.imgur.com/cGRAIwy.png)
