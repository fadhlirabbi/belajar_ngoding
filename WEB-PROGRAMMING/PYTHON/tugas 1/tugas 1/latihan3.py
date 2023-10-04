# import library
import time
import datetime

# input nama user
nama = input("Hallo... nama saya Mr. Kompie, nama Anda siapa? ")

# tampilkan nama user
print("Oh.. nama Anda", nama, ", nama yang bagus sekali.")

# kasih jeda 3 detik
time.sleep(3)

# input tahun lahir
thnLahir = int(input("BTW... " + nama + " kamu lahir tahun berapa? "))

# kasih jeda 3 detik
time.sleep(3)

# hitung usia user
skrg = datetime.datetime.now()
# usia = skrg.year + thnLahir                          # harusnya dikurang
usia = skrg.year - thnLahir

# tampilkan usia
print("Hmmm...", nama,"kamu sudah", usia,"tahun ya..") # variable nama pake x (namax)

# kasih jeda 3 detik
time.sleep(3)

# tampilkan pesan sesuai range usia
if (usia > 50):
    print("Anda sudah cukup tua ya?")                  # setelah tanda ? tidak pake tanda "
    print("Jaga kesehatan ya!!")

elif (usia > 20):                                      # kurang : setelah )
    print("Ternyata Anda masih cukup muda belia")
    print("Jangan sia-siakan masa mudamu ya!!")

elif (usia > 17):
    print("Hihihi... Anda ternyata masih ABG")
    print("Mulai berpikirlah secara dewasa ya!!")      # tidak ada tanda ) di akhir

else:
    print("Oalah.. Anda masih anak-anak toh?")
    print("Jangan suka merengek-rengek minta jajan ya!!")

# kasih jeda 3 detik
time.sleep(3)

# say goodbye
print("OK.. see you later", nama, ".. senang berkenalan denganmu")

# ====================================================================================

"""
JAWABAN DARI SOAL TADI

7. input data
    rumus cek semua kondisi 
    usia 1 = 50 < x
    usia2 = 50 > x > 20 
    usia3 = 20 > x > 17
    usia4 = usia1 & usia2 & usia3 = false (tidak terpenuhi)

8. fix error tahun, seharusnya (-)
    menjadi --> usia = skrg.year - thnLahir

10. penjelasan kegunaan dari :

    A. Tanda # berfungsi untuk menambahkan komentar 
    B. Perintah input() berfungsi untuk membuat form input dari user terhadap program yang memerlukan   input dari penguna untuk menjalankan program
    C. Perintah int() adalah suatu tipe data yang mendeklarasikan suatu bilangan bulat saja

11. Apa yang terjadi jika perintah import time dihapus?
    yang terjadi adalah kita tidak bisa mendapatkan tahun pada saat sekarang. akibatnya kita tidak bisa menggunakan fungsi dari library "datetime.datetime.now()", sehingga variable skrg akan menjadi error


"""