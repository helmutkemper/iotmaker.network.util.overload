package iotmakernetworkutiloverload

// ParserFunc (English): Optional parser function.
//   This function place is between receive and send data package.
//     inData    - Received data from connection
//     inLength  - Received data length
//     direction - Received data direction
//     outData   - Received data after parser
//     outLength - Received data length after parser
//     err       - Error
//
// ParserFunc (Português): Função de parser opcional.
//   Este função fica localizada entre o receptor e o transmissor de pacote de dados.
//     inData    - Dado recebido pela conexão
//     inLength  - Tamanho do dado recebido
//     direction - Direção do dado recebido
//     outData   - Dado recebido depois do parser
//     outLength - Tamanho do dado recebido depois do parser
//     err       - Erro
type ParserFunc func(
	inData []byte,
	inLength int,
	direction Direction,
) (
	outData []byte,
	outLength int,
	err error,
)
