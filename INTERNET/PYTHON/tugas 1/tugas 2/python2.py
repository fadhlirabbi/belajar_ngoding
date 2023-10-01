# status boolean dari nilai kelulusan
boolstatus = None

statusValid = "LULUS"
statusNotValid ="TIDAK LULUS"
printStatus = ("\nSTATUS                       :")

# input 3 macam nilai
bindo = input("Masukkan Nilai BHS Indonesia : ")
ipa = input("Masukkan Nilai IPA           : ")
matematika = input("Masukkan Nilai Matematika    : ") 

# check tipe 3x data 
if (bindo != int):
     print(printStatus,"Nilai BHS Indonesia Tidak Berupa Angka")
if (ipa != int):
     print(printStatus,"Nilai IPA Tidak Berupa Angka")
if(matematika != int):
     print(printStatus,"Nilai Matematika Tidak Berupa Angka")

# cek range 1-100
if bindo not in range(101):
     print(printStatus,"Nilai BHS Indonesia Tidak Berada di Range 1 - 100")
if ipa not in range(101):
     print(printStatus,"Nilai IPA Tidak Berada di Range 1 - 100")
if matematika not in range(101):
     print(printStatus,"Nilai Matemartika Tidak Berada di Range 1 - 100")