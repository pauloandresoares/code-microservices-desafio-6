# Code Education - desafio 5 - Kubernetes e hpa

### 1) Implementar algoritimo em Go Lang (looping somando a raiz quadrada):

+ Nome do deployment, service e nome da imagem deverá se chamar: go-hpa
+ Testes
+ Processo de CI
+ Push da imagem no Docker Hub (seu-user/go-hpa)
+ Deploy da imagem no K8S (arquivos de deployment e services)
+ Cada réplica deverá consumir no mínimo 50m e no máximo 100m.

### 2) O "hpa" do deployment tem as seguintes características:

+ O processo de escala inicia quando a CPU passar de 15%
+ Quantidade mínima de pods: 1
+ Quantidade máxima de pods: 6

### 3) Fazer teste usando requisições através de um looping e analisar o autoscaler.

Entrega:

