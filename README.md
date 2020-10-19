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

### Use example
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

### Optional data dump for testing
```
in:
00000000  e9 00 00 00 05 00 00 00  00 00 00 00 d4 07 00 00  |................|
00000010  04 00 00 00 61 64 6d 69  6e 2e 24 63 6d 64 00 00  |....admin.$cmd..|
00000020  00 00 00 ff ff ff ff c2  00 00 00 10 69 73 4d 61  |............isMa|
00000030  73 74 65 72 00 01 00 00  00 04 63 6f 6d 70 72 65  |ster......compre|
00000040  73 73 69 6f 6e 00 05 00  00 00 00 03 63 6c 69 65  |ssion.......clie|
00000050  6e 74 00 95 00 00 00 03  64 72 69 76 65 72 00 3e  |nt......driver.>|
00000060  00 00 00 02 6e 61 6d 65  00 10 00 00 00 6d 6f 6e  |....name.....mon|
00000070  67 6f 2d 67 6f 2d 64 72  69 76 65 72 00 02 76 65  |go-go-driver..ve|
00000080  72 73 69 6f 6e 00 12 00  00 00 76 31 2e 34 2e 30  |rsion.....v1.4.0|
00000090  2b 70 72 65 72 65 6c 65  61 73 65 00 00 03 6f 73  |+prerelease...os|
000000a0  00 2f 00 00 00 02 74 79  70 65 00 08 00 00 00 77  |./....type.....w|
000000b0  69 6e 64 6f 77 73 00 02  61 72 63 68 69 74 65 63  |indows..architec|
000000c0  74 75 72 65 00 06 00 00  00 61 6d 64 36 34 00 00  |ture.....amd64..|
000000d0  02 70 6c 61 74 66 6f 72  6d 00 09 00 00 00 67 6f  |.platform.....go|
000000e0  31 2e 31 34 2e 33 00 00  00                       |1.14.3...|

out:
00000000  3f 01 00 00 bb 0e 00 00  05 00 00 00 01 00 00 00  |?...............|
00000010  08 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000020  01 00 00 00 1b 01 00 00  08 69 73 6d 61 73 74 65  |.........ismaste|
00000030  72 00 01 03 74 6f 70 6f  6c 6f 67 79 56 65 72 73  |r...topologyVers|
00000040  69 6f 6e 00 2d 00 00 00  07 70 72 6f 63 65 73 73  |ion.-....process|
00000050  49 64 00 5f 8a fc 0e df  82 11 f7 45 b8 aa b9 12  |Id._.......E....|
00000060  63 6f 75 6e 74 65 72 00  00 00 00 00 00 00 00 00  |counter.........|
00000070  00 10 6d 61 78 42 73 6f  6e 4f 62 6a 65 63 74 53  |..maxBsonObjectS|
00000080  69 7a 65 00 00 00 00 01  10 6d 61 78 4d 65 73 73  |ize......maxMess|
00000090  61 67 65 53 69 7a 65 42  79 74 65 73 00 00 6c dc  |ageSizeBytes..l.|
000000a0  02 10 6d 61 78 57 72 69  74 65 42 61 74 63 68 53  |..maxWriteBatchS|
000000b0  69 7a 65 00 a0 86 01 00  09 6c 6f 63 61 6c 54 69  |ize......localTi|
000000c0  6d 65 00 25 13 d0 3e 75  01 00 00 10 6c 6f 67 69  |me.%..>u....logi|
000000d0  63 61 6c 53 65 73 73 69  6f 6e 54 69 6d 65 6f 75  |calSessionTimeou|
000000e0  74 4d 69 6e 75 74 65 73  00 1e 00 00 00 10 63 6f  |tMinutes......co|
000000f0  6e 6e 65 63 74 69 6f 6e  49 64 00 de 01 00 00 10  |nnectionId......|
00000100  6d 69 6e 57 69 72 65 56  65 72 73 69 6f 6e 00 00  |minWireVersion..|
00000110  00 00 00 10 6d 61 78 57  69 72 65 56 65 72 73 69  |....maxWireVersi|
00000120  6f 6e 00 09 00 00 00 08  72 65 61 64 4f 6e 6c 79  |on......readOnly|
00000130  00 00 01 6f 6b 00 00 00  00 00 00 00 f0 3f 00     |...ok........?.|

in:
00000000  37 00 00 00 06 00 00 00  00 00 00 00 dd 07 00 00  |7...............|
00000010  00 00 00 00 00 22 00 00  00 10 69 73 4d 61 73 74  |....."....isMast|
00000020  65 72 00 01 00 00 00 02  24 64 62 00 06 00 00 00  |er......$db.....|
00000030  61 64 6d 69 6e 00 00                              |admin..|

in:
00000000  e9 00 00 00 07 00 00 00  00 00 00 00 d4 07 00 00  |................|
00000010  04 00 00 00 61 64 6d 69  6e 2e 24 63 6d 64 00 00  |....admin.$cmd..|
00000020  00 00 00 ff ff ff ff c2  00 00 00 10 69 73 4d 61  |............isMa|
00000030  73 74 65 72 00 01 00 00  00 04 63 6f 6d 70 72 65  |ster......compre|
00000040  73 73 69 6f 6e 00 05 00  00 00 00 03 63 6c 69 65  |ssion.......clie|
00000050  6e 74 00 95 00 00 00 03  64 72 69 76 65 72 00 3e  |nt......driver.>|
00000060  00 00 00 02 6e 61 6d 65  00 10 00 00 00 6d 6f 6e  |....name.....mon|
00000070  67 6f 2d 67 6f 2d 64 72  69 76 65 72 00 02 76 65  |go-go-driver..ve|
00000080  72 73 69 6f 6e 00 12 00  00 00 76 31 2e 34 2e 30  |rsion.....v1.4.0|
00000090  2b 70 72 65 72 65 6c 65  61 73 65 00 00 03 6f 73  |+prerelease...os|
000000a0  00 2f 00 00 00 02 74 79  70 65 00 08 00 00 00 77  |./....type.....w|
000000b0  69 6e 64 6f 77 73 00 02  61 72 63 68 69 74 65 63  |indows..architec|
000000c0  74 75 72 65 00 06 00 00  00 61 6d 64 36 34 00 00  |ture.....amd64..|
000000d0  02 70 6c 61 74 66 6f 72  6d 00 09 00 00 00 67 6f  |.platform.....go|
000000e0  31 2e 31 34 2e 33 00 00  00                       |1.14.3...|

out:
00000000  30 01 00 00 bc 0e 00 00  06 00 00 00 dd 07 00 00  |0...............|
00000010  00 00 00 00 00 1b 01 00  00 08 69 73 6d 61 73 74  |..........ismast|
00000020  65 72 00 01 03 74 6f 70  6f 6c 6f 67 79 56 65 72  |er...topologyVer|
00000030  73 69 6f 6e 00 2d 00 00  00 07 70 72 6f 63 65 73  |sion.-....proces|
00000040  73 49 64 00 5f 8a fc 0e  df 82 11 f7 45 b8 aa b9  |sId._.......E...|
00000050  12 63 6f 75 6e 74 65 72  00 00 00 00 00 00 00 00  |.counter........|
00000060  00 00 10 6d 61 78 42 73  6f 6e 4f 62 6a 65 63 74  |...maxBsonObject|
00000070  53 69 7a 65 00 00 00 00  01 10 6d 61 78 4d 65 73  |Size......maxMes|
00000080  73 61 67 65 53 69 7a 65  42 79 74 65 73 00 00 6c  |sageSizeBytes..l|
00000090  dc 02 10 6d 61 78 57 72  69 74 65 42 61 74 63 68  |...maxWriteBatch|
000000a0  53 69 7a 65 00 a0 86 01  00 09 6c 6f 63 61 6c 54  |Size......localT|
000000b0  69 6d 65 00 ad 26 d0 3e  75 01 00 00 10 6c 6f 67  |ime..&.>u....log|
000000c0  69 63 61 6c 53 65 73 73  69 6f 6e 54 69 6d 65 6f  |icalSessionTimeo|
000000d0  75 74 4d 69 6e 75 74 65  73 00 1e 00 00 00 10 63  |utMinutes......c|
000000e0  6f 6e 6e 65 63 74 69 6f  6e 49 64 00 de 01 00 00  |onnectionId.....|
000000f0  10 6d 69 6e 57 69 72 65  56 65 72 73 69 6f 6e 00  |.minWireVersion.|
00000100  00 00 00 00 10 6d 61 78  57 69 72 65 56 65 72 73  |.....maxWireVers|
00000110  69 6f 6e 00 09 00 00 00  08 72 65 61 64 4f 6e 6c  |ion......readOnl|
00000120  79 00 00 01 6f 6b 00 00  00 00 00 00 00 f0 3f 00  |y...ok........?.|

out:
00000000  3f 01 00 00 bd 0e 00 00  07 00 00 00 01 00 00 00  |?...............|
00000010  08 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000020  01 00 00 00 1b 01 00 00  08 69 73 6d 61 73 74 65  |.........ismaste|
00000030  72 00 01 03 74 6f 70 6f  6c 6f 67 79 56 65 72 73  |r...topologyVers|
00000040  69 6f 6e 00 2d 00 00 00  07 70 72 6f 63 65 73 73  |ion.-....process|
00000050  49 64 00 5f 8a fc 0e df  82 11 f7 45 b8 aa b9 12  |Id._.......E....|
00000060  63 6f 75 6e 74 65 72 00  00 00 00 00 00 00 00 00  |counter.........|
00000070  00 10 6d 61 78 42 73 6f  6e 4f 62 6a 65 63 74 53  |..maxBsonObjectS|
00000080  69 7a 65 00 00 00 00 01  10 6d 61 78 4d 65 73 73  |ize......maxMess|
00000090  61 67 65 53 69 7a 65 42  79 74 65 73 00 00 6c dc  |ageSizeBytes..l.|
000000a0  02 10 6d 61 78 57 72 69  74 65 42 61 74 63 68 53  |..maxWriteBatchS|
000000b0  69 7a 65 00 a0 86 01 00  09 6c 6f 63 61 6c 54 69  |ize......localTi|
000000c0  6d 65 00 b0 26 d0 3e 75  01 00 00 10 6c 6f 67 69  |me..&.>u....logi|
000000d0  63 61 6c 53 65 73 73 69  6f 6e 54 69 6d 65 6f 75  |calSessionTimeou|
000000e0  74 4d 69 6e 75 74 65 73  00 1e 00 00 00 10 63 6f  |tMinutes......co|
000000f0  6e 6e 65 63 74 69 6f 6e  49 64 00 df 01 00 00 10  |nnectionId......|
00000100  6d 69 6e 57 69 72 65 56  65 72 73 69 6f 6e 00 00  |minWireVersion..|
00000110  00 00 00 10 6d 61 78 57  69 72 65 56 65 72 73 69  |....maxWireVersi|
00000120  6f 6e 00 09 00 00 00 08  72 65 61 64 4f 6e 6c 79  |on......readOnly|
00000130  00 00 01 6f 6b 00 00 00  00 00 00 00 f0 3f 00     |...ok........?.|

in:
00000000  57 00 00 00 08 00 00 00  00 00 00 00 dd 07 00 00  |W...............|
00000010  00 00 00 00 00 42 00 00  00 10 70 69 6e 67 00 01  |.....B....ping..|
00000020  00 00 00 03 6c 73 69 64  00 1e 00 00 00 05 69 64  |....lsid......id|
00000030  00 10 00 00 00 04 ac a1  10 37 e6 e1 40 8a a4 b6  |.........7..@...|
00000040  5d 87 06 5a 45 a6 00 02  24 64 62 00 06 00 00 00  |]..ZE...$db.....|
00000050  61 64 6d 69 6e 00 00                              |admin..|

in:
00000000  be 00 00 00 09 00 00 00  00 00 00 00 dd 07 00 00  |................|
00000010  00 00 00 00 00 53 00 00  00 02 69 6e 73 65 72 74  |.....S....insert|
00000020  00 06 00 00 00 64 69 6e  6f 73 00 08 6f 72 64 65  |.....dinos..orde|
00000030  72 65 64 00 01 03 6c 73  69 64 00 1e 00 00 00 05  |red...lsid......|
00000040  69 64 00 10 00 00 00 04  ac a1 10 37 e6 e1 40 8a  |id.........7..@.|
00000050  a4 b6 5d 87 06 5a 45 a6  00 02 24 64 62 00 05 00  |..]..ZE...$db...|
00000060  00 00 74 65 73 74 00 00  01 55 00 00 00 64 6f 63  |..test...U...doc|
00000070  75 6d 65 6e 74 73 00 47  00 00 00 07 5f 69 64 00  |uments.G...._id.|
00000080  5f 8d 02 1b 0a 58 9d 35  80 13 27 0d 02 6e 61 6d  |_....X.5..'..nam|
00000090  65 00 06 00 00 00 54 2d  52 65 78 00 10 61 67 65  |e.....T-Rex..age|
000000a0  00 0a 00 00 00 02 63 69  74 79 00 0e 00 00 00 4a  |......city.....J|
000000b0  75 72 61 73 73 69 63 20  54 6f 77 6e 00 00        |urassic Town..|

out:
00000000  26 00 00 00 be 0e 00 00  08 00 00 00 dd 07 00 00  |&...............|
00000010  00 00 00 00 00 11 00 00  00 01 6f 6b 00 00 00 00  |..........ok....|
00000020  00 00 00 f0 3f 00                                 |....?.|

out:
00000000  2d 00 00 00 bf 0e 00 00  09 00 00 00 dd 07 00 00  |-...............|
00000010  00 00 00 00 00 18 00 00  00 10 6e 00 01 00 00 00  |..........n.....|
00000020  01 6f 6b 00 00 00 00 00  00 00 f0 3f 00           |.ok........?.|

End!
Duration: 15.0030159s
```