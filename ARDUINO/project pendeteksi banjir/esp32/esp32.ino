#include <Wire.h>

// library wifi
#include "WiFi.h"

// library webserver
#include "HTTPClient.h"

// define ssid and pass wifi
const char* ssid = "ytta";
const char* pass = "yttaytta";

// variable host server yang akan menampung database dan aplikasi web
// 4F = 192.168.100.6
// laptop = 192.168.84.41
// kalau akai xampp pakai ip, kalo pake hosting pakai nama webnya monitoringonlineketinggianair.000webhostapp.com
const char* host = "192.168.178.41";

char a;
String terima;
int realdistance;

int terimaInt;
int index1, index2;
String dataSubs;

void setup(){
  Serial.begin (9600);
  
  delay(4000);
  WiFi.begin(ssid, pass);
  Serial.print("connecting...");

  while(WiFi.status() != WL_CONNECTED){

    Serial.print(".");
    delay(1000);
  }

  // jika terdeteksi 
  Serial.println("Connected");
}

void loop(){

  
  if (Serial.available()){
    a = Serial.read();
    terima += a;

    if (terima.length()>0 && a == '0'){
      index1 = terima.indexOf('!');
      index2 = terima.indexOf('0', index1 + 1);

      dataSubs = terima.substring(index1+1, index2);

      realdistance = dataSubs.toInt();
      terima = "";

    }
  }

  //Serial.println("ketinggian air : " + int (realdistance));
  Serial.print(realdistance);
  Serial.print(" cm\n");
  //delay(1000);

  
  // kirim data ke server
  WiFiClient client;

  // inisialisasi port web server (80)
  const int httpPort = 80;

  if(!client.connect(host, httpPort)){

    Serial.println("Connection Failed");
    return;
  }

  // kondisi terpenuhi ---> kirim database ke webserver
  String Link;
  HTTPClient http;

  // kalau pakai xampp
  Link = "http://" + String(host) + "/webserverbanjir/kirimdata.php?sensor=" + String(realdistance);
  
  // kalau pakai hosting nama database nya dihapus
  //Link = "http://" + String(host) + "/kirimdata.php?sensor=" + String(realdistance);
  // eksekusi alamat Link
  http.begin(Link);
  http.GET();  

  
  //baca respon setelah berhasil kirim nilai sensor
  String respon = http.getString();
  Serial.println(respon);
  http.end();
  
}