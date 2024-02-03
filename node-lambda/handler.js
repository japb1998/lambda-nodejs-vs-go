'use strict';
let s3 = require('@aws-sdk/client-s3')

let client = new s3.S3Client();

const hello = async (event) => {

  const command = new s3.GetObjectCommand({
    Bucket: process.env.BUCKET,
    Key: "hello.txt",
  });

  let str = "";

  try {
    const response = await client.send(command);
    // The Body object also has 'transformToByteArray' and 'transformToWebStream' methods.
    str = await response.Body.transformToString();

  } catch (err) {
   
    return {
      statusCode: 500,
      body: JSON.stringify(err.message),
    }
  }
  return {
    statusCode: 200,
    body: str,
  };

  // Use this code if you don't use the http event with the LAMBDA-PROXY integration
  // return { message: 'Go Serverless v1.0! Your function executed successfully!', event };
};
module.exports.hello = hello;