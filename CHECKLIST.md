# check list 
- common/const.go Version
- CHANGELOG.md
- Makefile
- merge branch
- docker build and push to docker hub(//todo nick: auto build and auto tag )
  - make build-doctron-alpine
  - docker tag lampnick/doctron:latest lampnick/doctron:v0.3.2-alpine
  - docker tag lampnick/doctron:latest lampnick/html2pdf:latest
  - docker tag lampnick/doctron:latest lampnick/html2pdf:v0.3.2-alpine
  - docker tag lampnick/doctron:latest lampnick/html2image:latest
  - docker tag lampnick/doctron:latest lampnick/html2image:v0.3.2-alpine
  - docker push lampnick/doctron:latest
  - docker push lampnick/doctron:v0.3.2-alpine
  - docker push lampnick/html2pdf:latest
  - docker push lampnick/html2pdf:v0.3.2-alpine
  - docker push lampnick/html2image:latest
  - docker push lampnick/html2image:v0.3.2-alpine
