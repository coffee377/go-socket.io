const {encodePacket,encodePayload} = require('engine.io-parser')

// 72 101 108 108 111 44 32 87 111 114 108 100 33
const data  = Buffer.from('Hello, World!', 'utf8')
encodePacket({type: 'message', data}, true, function (encodedValue) {
    console.log(encodedValue)
    // for (let i = 0; i < encodedValue.length; i++) {
    //     console.log(encodedValue[i])
    // }
})


const inputMessages = ["Engine.IO",Buffer.from("Test.Data", 'utf8')];
const packets = [];
for (let i = 0; i < inputMessages.length; i++) {
    packets.push({ type: 'message', data: inputMessages[i] });
}

encodePayload(packets, function (encodedValue) {
    console.log(encodedValue)
});