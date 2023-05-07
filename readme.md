## ubm - Unity Build Manager
Unity Gaming Services provides Unity game developers with the ability to create builds of their games.

### Problems

Unity is often picked by developers for cross-platform build convenience.<br>
Creating a traditional pipeline to handle different builds for different platforms is long and complicated.<br>
Unity Gaming Services simplifies this by making it as simple to auto build as it is to build locally.<br>
Unfortunately there is little support for post build steps, for those that want to create deployment steps etc you are limited to [Script Hooks](https://unity-technologies.github.io/cloud-build-docs-public/windows-builder-scripts) and trying to connect these to 3rd party ci/cd for better support.<br>
 

### How ubm solves these problems
Using the [Build Api](https://build-api.cloud.unity3d.com/docs/1.0.0/index.html), you can access most of the features in the Unity Gaming Services.
- By first creating a "build target" in Unity Gaming Services
- ubm triggers a new build for this "build target" via the api. 
- ubm waits for this build to complete.
- once complete the api adds the storage location of the build artifact to the return object
- ubm then downloads the artifact
- if run within a third party ci/cd you are now able to do as you wish with the artifact
  for example: deploy the webgl version onto staging/production server. 


### Security
Unity handles access to its own `https://build-api.cloud.unity3d.com/api/v1/` using an api key you can access from the Unity Gaming Services console.<br>
In order to get the artifact url you must authenticate with the Unity api however unfortunately the Google storage location of the artifacts are not secured. <br>
This leads me to believe you could brute force build artifact locations but haven't tried this.<br>
Unity has not included the build api or is storage domains in their [Bug Bounty](https://bugcrowd.com/unity) though will accept out of scope domains (just wont pay ya).<br>