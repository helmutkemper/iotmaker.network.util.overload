# iotmaker.network.util.overload

Simulates a slow network between two connections
> Work in progress

```

 Normal use: create and insert data, ~30ms
 +------------------+                                  +------------------+
 |                  |                                  |                  |
 |  Golang MongoDB  | -------------------------------> | MongoDB Database |
 |      Driver      |                                  |                  |
 |    Port 27017    | <------------------------------- |    Port 27017    |
 |                  |                                  |                  |
 +------------------+                                  +------------------+

 Proposed use: create and insert data, ~5s/~15s
 +------------------+       +------------------+       +------------------+
 |                  |       |                  |       |                  |
 |  Golang MongoDB  | ----> | Network Overload | ----> | MongoDB Database |
 |      Driver      |       | IN:   Port 27016 |       |                  |
 |    Port 27016    | <---- | OUT:  Port 27017 | <---- |    Port 27017    |
 |                  |       |                  |       |                  |
 +------------------+       +------------------+       +------------------+

```

### English:

This project simulates an overloaded network by inserting waits between network packets, 
in an attempt to simulate a network much slower than the conditions normally found on 
local networks.

The main objective of this project is to detect design problems in the source codes in 
relation to the network times found in production, however, difficult to simulate in the 
development environment.

### Português:

Este projeto simula uma rede sobrecarregada inserindo esperas entre os pacotes de rede, 
na tentativa de simular uma rede muito mais lenta do que as condições encontradas 
normalmente nas redes locais.

O principal objetivo desse projeto é detectar problemas de concepção nos código fonte em 
relação aos tempos de rede encontrados em produção, porém, de difícil simulação no 
ambiente de desenvolvimento.

```golang
package main

import (
	"context"
	"encoding/hex"
	"fmt"
	overload "github.com/helmutkemper/iotmaker.network.util.overload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"runtime/debug"
	"time"
)

// main (English): Sample program. Creates a database named 'test' and a collection named
// 'dino' in a local MongoDB address '127.0.0.1:27017' in a time between 1.5 seconds
// and 15 seconds, considering a local machine without much process request .
//
// main (Português): Programa de exemplo. Cria um banco de dados de nome 'test' e uma
// coleção de nome 'dino' em um MongoDB local de endereço '127.0.0.1:27017' em um
// tempo entre 1.5 segundos e 15 segundos, considerando uma máquina local sem muita
// requisição de processos.
func main() {
	var err error

	// (English): MongoDB timeout
	// (Português): Tempo limite dos processos do mongoDB
	var timeout = time.Millisecond * 1000 * 30

  // (English): Minimal delay between packages, 0.5 seconds
  // (Português): Atraso mínimo inserido entre os pacotes, 0.5 segundos
  var delayMin = time.Millisecond * 500

  // (English): Maximal delay between packages, 5 seconds
  // (Português): Atraso máximo inserido entre os pacotes, 5 segundos
  var delayMax = time.Millisecond * 5000

  // (English): Test a local MongoDB connection
  // (Português): Testa a conexão com o MongoDB local
  err = testNormalMongoDB(
	  "mongodb://127.0.0.1:27017",
	  timeout,
  )
	if err != nil {
		panic(string(debug.Stack()))
	}

  // (English): Prepares to divert port 27017 to 27016.
  // Note: Every connection has a direction. In the case of a database, the connection is
  // from the driver to the bank, so the input address must be the driver address and the
  // bank address must be the output address
  //
  // (Português): Prepara o desvio da porta 27017 para 27016.
  // Entenda: Toda conexão tem um sentido. No caso de banco de dados, a conexão é do
  // driver para o banco, então, o endereço de entrada deve ser o endereço do driver e o
  // endereço do banco deve ser o endereço de saída
  err = mountNetworkOverload(
	  delayMin,
	  delayMax,
	  "127.0.0.1:27016",
	  "127.0.0.1:27017",
  )
	if err != nil {
		panic(string(debug.Stack()))
	}

  // (English): Tenta criar um novo banco de dados e coleção usando uma conexão de dados
  // mais lenta da porta 27016 para que possa haver interferência no caminho dos pacotes.
  //
  // (Português): Tenta criar um novo banco de dados e coleção usando uma conexão de
  // dados mais lenta da porta 27016 para que possa haver interferência no caminho dos
  // pacotes.
  err = testNetworkOverloaded(
	  "mongodb://127.0.0.1:27016",
	  timeout,
  )
	if err != nil {
		panic(string(debug.Stack()))
	}
}

// binaryDump (English): Custom function, used to interfere in the data, in case there is
// any need for processing.
//
// binaryDump (Português): Função customizada, usada interferir nos dados, caso haja
// alguma necessidade de processamento.
func binaryDump(
  inData []byte,
  inLength int,
  direction overload.Direction,
) (
  outData []byte,
  outLength int,
  err error,
) {

  // (English): Copy the input data to the output data.
  // (Português): Copia o dado de entrada no dado de saída.
  outData = inData

  // (English): Copy the sizes of the data buffers.
  // (Português): Copia os tamanhos dos buffers de dados.
  outLength = inLength

  // (English): Prints the direction of the data. input is from the drive to the bank and
  // output is from the bank to the drive.
  // (Português): Imprime a direção do dado. entrada é do drive para o banco e saída é do
  // banco para o drive.
  fmt.Printf("%v:\n", direction)

  // (English): Human buffer prints data
  // (Português): Imprime o dado do buffer de forma humana
  fmt.Printf("%v\n", hex.Dump(inData[:inLength]))

	return
}

// mountNetworkOverload (English): Assemble the proxy with the data of the new connection
// mountNetworkOverload (Português): Monta o proxy com os dados da nova conexão
func mountNetworkOverload(
  min time.Duration,
  max time.Duration,
  inAddress string,
  outAddress string,
) (
  err error,
) {

  // (English): Prepare the driver for TCP network
  // (Português): Prepara o driver para rede TCP
  var over = &overload.NetworkOverload{
		ConnectionInterface: &overload.TCPConnection{},
	}

  // (English): Enables the TCP protocol and the input and output addresses
  // (Português): Habilita o protocolo TCP e os endereços de entrada e saída
  err = over.SetAddress(overload.KTypeNetworkTcp, inAddress, outAddress)
	if err != nil {
		return
	}

  // (English): [optional] Points to the custom function for data processing
  // (Português): [opcional] Aponta a função personalizada para tratamento dos dados
  over.ParserAppendTo(binaryDump)

  // (English): Determines the maximum and minimum times between packages
  // (Português): Determina os tempos máximo e mínimos entre os pacotes
  over.SetDelay(min, max)

  // (English): Listen to port 27016 without blocking the code
  // (Português): Escuta a porta 27016 sem bloquear o código
  go func() {
		err = over.Listen()
		if err != nil {
			panic(string(debug.Stack()))
		}
	}()

	return
}

// testNormalMongoDB (English): Test local MongoDB to make sure it's working
// testNormalMongoDB (Português): Testa o MongoDB local para garantir que está
// funcionando
func testNormalMongoDB(
  address string,
  timeOut time.Duration,
) (
  err error,
) {

	var mongoClient *mongo.Client
	var ctx context.Context

  // (English): Prepare the MongoDB client
  // (Português): Prepara o cliente do MongoDB
  mongoClient, err = mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		return
	}

  // (English): Prepares the timeout
  // (Português): Prepara o tempo limite
  ctx, _ = context.WithTimeout(context.Background(), timeOut)

  // (English): Connects to MongoDB
  // (Português): Conecta ao MongoDB
  err = mongoClient.Connect(ctx)
	if err != nil {
		return
	}

	var cancel context.CancelFunc

  // (English): Ping() to ensure local MongoDB is working before testing
  // (Português): Faz um ping() para garantir que o MongoDB local está funcionando antes
  // do teste
  ctx, cancel = context.WithTimeout(context.Background(), timeOut)
  err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}
	defer cancel()

  // (English): Disconnects from the bank at the end of the test
  // (Português): Desconecta do banco ao final do teste
  err = mongoClient.Disconnect(ctx)

	return
}

// testNetworkOverloaded (English): Tests the new network port
// testNetworkOverloaded (Português): Testa a nova porta de rede
func testNetworkOverloaded(
  address string,
  timeout time.Duration,
) (
  err error,
) {

  // (English): Runtime measurement starts
  // (Português): Começa a medição do tempo de execução
  start := time.Now()

	var mongoClient *mongo.Client
	var cancel context.CancelFunc
	var ctx context.Context

  // (English): Prepare the MongoDB client
  // (Português): Prepara o cliente do MongoDB
  mongoClient, err = mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		return
	}

  // (English): Connects to MongoDB
  // (Português): Conecta ao MongoDB
  err = mongoClient.Connect(ctx)
	if err != nil {
		return
	}

  // (English): Prepares the timeout
  // (Português): Prepara o tempo limite
  ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

  // (English): Ping() to test the MongoDB connection
  // (Português): Faz um ping() para testar a conexão do MongoDB
  err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

  // (English): New collection format
	// (Português): Formato da nova coleção
	type Trainer struct {
		Name string
		Age  int
		City string
	}

  // (English): Creates the 'test' bank and the 'dinos' collection
  // (Português): Cria o banco 'test' e a coleção 'dinos'
  collection := mongoClient.Database("test").Collection("dinos")

  // (English): Prepares the data to be inserted
  // (Português): Prepara os dados a serem inseridos
  trainerData := Trainer{"T-Rex", 10, "Jurassic Town"}

  // (English): Insert the data
  // (Português): Insere os dados
  _, err = collection.InsertOne(context.TODO(), trainerData)
	if err != nil {
		return
	}

  // (English): Stop the operation time measurement
  // (Português): Para a medição de tempo da operação
  duration := time.Since(start)
  fmt.Printf("End!\n")
  fmt.Printf("Duration: %v\n\n", duration)

	return
}
```