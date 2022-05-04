<div align="center">


  <img src="assets/icon-512x512.png" alt="logo" width="200" height="auto" />
  <h1>Cocopasty</h1>
  
  <p>
    A simple copy-and-paste app to selfhost
  </p>
  
  <p>
    ⚠️ This project might still have a few bugs and things to test ⚠️ 
  </p>
  
  
<!-- Badges -->
<p>
  <a href="https://github.com/MasterEvarior/cocopasty/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/MasterEvarior/cocopasty" alt="contributors" />
  </a>
  <a href="">
    <img src="https://img.shields.io/github/last-commit/MasterEvarior/cocopasty" alt="last update" />
  </a>
  <a href="https://github.com/MasterEvarior/cocopasty/network/members">
    <img src="https://img.shields.io/github/forks/MasterEvarior/cocopasty" alt="forks" />
  </a>
  <a href="https://github.com/MasterEvarior/cocopasty/stargazers">
    <img src="https://img.shields.io/github/stars/MasterEvarior/cocopasty" alt="stars" />
  </a>
  <a href="https://github.com/MasterEvarior/cocopasty/issues/">
    <img src="https://img.shields.io/github/issues/MasterEvarior/cocopasty" alt="open issues" />
  </a>
  <a href="https://github.com/MasterEvarior/cocopasty/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/MasterEvarior/cocopasty" alt="license" />
  </a>
</p>
   
<h4>
    <a href="https://github.com/MasterEvarior/cocopasty/issues/new?assignees=MasterEvarior&labels=bug&template=bug_report.md&title=">Report Bug</a>
  <span> · </span>
    <a href="https://github.com/MasterEvarior/cocopasty/issues/new?assignees=MasterEvarior&labels=enhancement&template=feature_request.md&title=">Request Feature</a>
  </h4>
</div>

<br />

<!-- Table of Contents -->
# Table of Contents

- [About the Project](#about-the-project)
  * [Screenshots](#screenshots)
  * [Tech Stack](#tech-stack)
  * [Features](#features)
- [Getting Started](#getting-started)
  * [Configuration](#configuration)
  * [Prerequisites](#prerequisites)
  * [Run Locally](#run-locally)
  * [Deployment](#deployment)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [FAQ](#faq)
- [License](#license)
- [Contact](#contact)
- [Acknowledgements](#acknowledgements)
  

<!-- About the Project -->
## About the Project
Cocopasty is a small and simply copy-and-paste solution for your code. Paste your code in your browser and copy it on another device!

<!-- Screenshots -->
### Screenshots

<div align="center"> 
  <img src="assets/screenshot.png" alt="screenshot" />
</div>


<!-- TechStack -->
### Tech Stack

<details>
  <summary>Client</summary>
  <ul>
    <li><a href="https://www.vuejs.org/">Vue.js</a></li>
  </ul>
</details>

<details>
  <summary>Server</summary>
  <ul>
    <li><a href="https://go.dev/">Golang</a></li>
  </ul>
</details>

<details>
<summary>Database</summary>
  <ul>
    <li><a href="https://redis.io/">Redis</a></li>
  </ul>
</details>

<details>
<summary>DevOps</summary>
  <ul>
    <li><a href="https://www.docker.com/">Docker</a></li>
  </ul>
</details>

<!-- Features -->
### Features

- So simple, your pet rock could use it
- 150+ available languages
- Autodetect languages with [highlight.js](https://highlightjs.org)
- Completely dockerized
<!-- Getting Started -->
## Getting Started

<!-- Prerequisites -->
### Prerequisites
You need to have Go, Docker, Docker Compose and NPM installed to run/develop this project.

### Configuration
Explanations for the different environment variables. Examples can be found in the docker-compose.yml files.

#### Frontend
- `VUE_APP_BACKEND_PORT` is the port your cocopasty-backend container runs.
- `VUE_APP_BACKEND_HOST` is the host of your cocopasty-backend container.

#### Backend
- `LOG_LEVEL` the log level you wish to use. Default is "info". All log levels can be found [here](https://github.com/Sirupsen/logrus#level-logging).
- `REDIS_HOST` is the host + port of your Redis DB.

<!-- Run Locally -->
### Run Locally

Clone the project

```bash
  git clone https://https://github.com/MasterEvarior/cocopasty.git
```

Go to the project directory

```bash
  cd cocopasty
```

Run with Docker Compose

```bash
  docker-compose up
```

<!-- Deployment -->
### Deployment

#### Docker Compose
To deploy this project copy the contents of the [docker-compose-deploy.yml](https://github.com/MasterEvarior/cocopasty/blob/main/docker-compose-deploy.yml) file and paste it into a new file.

```bash
  wget https://raw.githubusercontent.com/MasterEvarior/cocopasty/main/docker-compose-deploy.yml
```
Change the ports, volumes, etc. to your liking.
```bash
  vi docker-compose-deploy.yml
```

Rename the file and run it with Docker Compose.
```bash
  mv docker-compose-deploy.yml docker-compose.yml
  docker-compose up
```
<!-- Roadmap -->
## Roadmap
Feel free to suggest features through a GitHub issue, in addition to the ones listed below:
- [ ] Backend tests
- [ ] Installation instructions for Unraid
- [ ] Installation instructions without Docker Compose

<!-- Contributing -->
## Contributing
Contributions are always welcome!

<!-- FAQ -->
## FAQ

- Is there an official method to install WITHOUT Docker?

  + No

- Is there any sort of build in authentication?

  + No, though you could use something like [Authelia](https://www.authelia.com/docs/)

- Can I use this to copy my API-Keys etc. from one device to another?

  + You can but I'd strongly advise against it

- Which languages are available for auto-highligting?

  + A complete list can be found [here](https://highlightjs.readthedocs.io/en/latest/supported-languages.html)

<!-- License -->
## License

Distributed under the MIT License. See the LICENSE file for more information.

<!-- Contact -->
## Contact

Email: contact@giannin.dev

Project Link: [https://github.com/MasterEvarior/cocopasty](https://github.com/MasterEvarior/cocopasty)


<!-- Acknowledgments -->
## Acknowledgements

Cool projects which make Cocopasty possible

 - [Shields.io](https://shields.io/)
 - [Awesome Readme Template](https://github.com/Louis3797/awesome-readme-template)
 - [highlight.js](https://highlightjs.org)
 - [Vue Prism Editor](https://github.com/koca/vue-prism-editor)
 - [Guide to Vue.js + Docker env variables](https://medium.com/js-dojo/vue-js-runtime-environment-variables-807fa8f68665)
 - [CSS Button Generator](https://www.bestcssbuttongenerator.com)
