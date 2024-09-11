# TEMPERATURE CONVERTER (C to F/R/K)

Riyan Naffa Nusafara

## Masukan Temperatur

Bagian kode pertama yaitu bagian untuk memberi _prompt_ masukan pada _terminal_ dengan mencetak `Temperature (in C): `.
Suatu variabel `temp_c` bertipe `float64` dideklarasi untuk kemudian diberi masukan.

**Kode:**
```go
fmt.Print("Temperature (in C): ")
var temp_c float64
fmt.Scan(&temp_c)
```
**Terminal:**
```
========== TEMPERATURE CONVERTER ==========

Temperature (in C):
```

## Masukan Opsi

Setelah masukan temperatur dalam celsius telah dimasukkan, _prompt_ opsi untuk ketiga konversi suhu ditampilkan.
Variabel `opt` bertipe `int` digunakan untuk menampung masukan opsi.

**Kode:**
```go
fmt.Println("Convert to:\n1. Fahrenheit\n2. Reamur\n3. Kelvin")
var opt int
fmt.Scan((&opt))
```

**Terminal:**
```go
Convert to:
1. Fahrenheit
2. Reamur
3. Kelvin

```

## Seleksi Opsi

Terdapat tiga opsi untuk masing-masing konversi suhu. Opsi 1 berarti konversi suhu dari celsius ke fahrenheit.
Opsi 2 berarti konversi suhu dari celsius ke reamur. Opsi 3 berarti konversi suhu dari celsius ke kelvin. Opsi _default_ berarti opsi tidak bernilai 1, 2, atau 3, yang artinya opsi tidak valid.

**Kode:**
```go
switch opt {
case 1:
  {
    // fahrenheit conversion
  }
case 2:
  {
    // reamur conversion
  }
case 3:
  {
    // kelvin conversion
  }
default:
  // invalid option
}
```

## Detail Konversi Suhu

Setiap konversi suhu memiliki rumus-rumus tersendiri.
- Celsius -> Fahrenheit

  $$F=\dfrac{9C}{5}+32$$
  
  **Kode:**
  ```go
  temp_f := temp_c * 9 / 5 + 32
  fmt.Printf("The temperature is %f F", temp_f)
  ```

  **Terminal:**
  ```
  The temperature is 32.000000 F
  ```
- Celsius -> Reamur

  $$R=\dfrac{4C}{5}$$

  **Kode:**
  ```go
  temp_r := temp_c * 4 / 5
  fmt.Printf("The temperature is %f R", temp_r)
  ```

  **Terminal:**
  ```
  The temperature is 0.000000 R
  ```
- Celsius -> Kelvin

  $$K=C+273.15$$

  **Kode:**
  ```go
  temp_k := temp_c + 273.15
  fmt.Printf("The temperature is %f K", temp_k)
  ```

  **Terminal:**
  ```
  The temperature is 273.150000 K
  ```
- Invalid Option

  **Kode:**
  ```go
  fmt.Println("Invalid option")
  ```

  **Terminal:**
  ```
  Invalid option
  ```
