#![allow(unused_imports)]
#![allow(dead_code)]

use std::net::{TcpListener, TcpStream};
use std::io::{Read, Write};
use std::thread;

fn main() {
  let listener = TcpListener::bind("127.0.0.1:8080").unwrap();

  for stream in listener.incoming() {
    let stream = stream.unwrap();

    thread::spawn(|| {
      handle_connection(stream);
    });
  }
}

fn handle_connection(mut stream: TcpStream) {
  let mut buffer = [0; 1024];

  stream.read(&mut buffer).unwrap();

  let response = b"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello, Envoy!";

  stream.write(response).unwrap();
}
