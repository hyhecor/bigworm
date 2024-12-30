
const Crypto = require("crypto");


const data_to_json = function(data){ //문자열을 json으로 복호화.
    var decipher = Crypto.createDecipher("aes-256-cbc", "iinpro");
    return JSON.parse(decipher.update(data.toString(), "base64", "utf8")
        +decipher.final("utf8"));
};

const json_to_data = function(json) {
    var cipher = Crypto.createCipher("aes-256-cbc", "iinpro");
    return cipher.update(JSON.stringify(json), "utf8", "base64")+cipher.final("base64");
};





// let s =  json_to_data(json);

// console.log(s)
// console.log(data_to_json(s))

// const plain = "4D/NZjMzUQa5ueLlGACTYXHxhSoO1a5dA6sw8Q2GO0e/ZhLPNm5JgctbsYSxT8fQTXdYHK4Loe5AdIfek+sEaqNIkU8vrPeOdBSS9OcQdZAXrrHqcAH5H+umYZmQcUoa"
// console.log(data_to_json(plain))



const net = require('net');


const client = net.connect({host: "dbos.co.kr", port: 10091}, () => {

    
    var json = {
        sCommand : "CONNECT",
        sMainCode : "0",
        sSubCode : "999",
        sIp : "127.0.0.1",
        sOs : "linux",
    };

    let s = json_to_data(json);

    console.log(s)

    client.write(s, err => console.log)
})
.on("end", _=>console.log("명령 서버 종료")) //end -> close
.on("timeout", console.log)
.on("data", (data) => {
    console.log(data_to_json(data));
    this.end();
  })
.on("close", _=>{
    console.log("command server close");
});


// client.on('data', (data) => {
//     console.log(data_to_json(data));
//     client.end();
//   });

//   client.on('error', (err) => {
//     console.log(err);
//   });


