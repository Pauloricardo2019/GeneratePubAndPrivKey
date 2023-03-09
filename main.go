package main

// const PRIVATE_KEY = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBMm5hUk5ORWp0TnZCOW94T3BFY1BWc1Vic3oyQUJWQXE1akVVZ2pSbzVCTWQyeE84CnYvM3h6eGZkRDdBdm52cUNCUWpwNXEvYXovUVVXLzdkOVJGUUJJL0o1VzF2UVdYOHlTUTZaVnVWS0VuWWU2ZWQKVmZFd1k2QnkzUThaWmg1dDJyWkVzTmlKa0VvNElJNEVweU5zZldUM3pGSkY4cmFpRGVIek9FOTZUSnYxLzBhVgpQdXlldEhhejdLVWxneUJvZ20yQ05tNkpzZWh3WWM3clNGaEgzbm1sS0hheWJwMFRUUkYzMVBYa25ZOXpaMWtzCkFZNVByVTE4dTBEZU5SVmtNRVZqRXUzZ00vaXplM0xUeUJFRkFuQmxTT08wMlUzaWNrRjh0OFhTM2todmFoeTUKeWdBR0srUGdVY05uK1M0T3FXM01NaU8yVjR3KzVPUnBFT2I4UXdJREFRQUJBb0lCQVFDOXB2ejcreXVpVUwwYwpDa1lqTzBWSHNLb1JlcXh1OEYrdG9xSUpQdkpvUFdoSHplU29VQTdpRi9kakxUTEtNdk00NTdvVytFZGZES2U0CkRaUW1zamJqSHVqSThhNEorajN2YURJQjEvV3ZHaHlEak1sNzJIeUhvTHpXeDhybTZqZTREVStvcnZNTmZrRW4Ka2VMY0xsWjZyWEtwd0NyVlFGT0M5Y3F0OGVGY0t0WHl1R1lTSVdLQzJJQlpObGFNQ2J4bEg5WWhBb055bDQ0bQp4TFRGSmwrRk5sVW9yVWQ4dkNjUTlTanFvQ3RRcG5ZazliY0ZLbHFOZWhBT1Zjam9GbzlpQjI0VjRYcHVMYVFPCnZucEU4dU9yTUViNUNKQm1XQVQ5RWgwazFsc2lhblFzalBmRDEzd2xQdFV0L1AwMEdzSDE1LzdOZENrWks2TzAKQnZnaW9ZNVJBb0dCQU9LVjVWN2NLbnNaa0xSdVAzQUMxYTA5aHkveWxhcVgyNG5HSTA0eGZFK3NIVkRsZFQvYgo5eE1LdjJkeXpPT3dXZStxQzIrenlweFo1bjZRdmszalhpM1pYWE83ZlJSOGVOMlhndzZaNEdMQVUxYWJRY3ROCmx5bkxjMEZVWG1zZTdmTDNQMEw2RnBxZ01VaC9IaUM2aUtrUXo4T1Q4ZW1qaVJyR1BKZE54d2RIQW9HQkFQYlMKdlo4Zi91ejg2YjdNd2xrR0ZPVjdBSFZYRHJyR3F4cHUxMjdLQXBCa2xmZWc1cmdlV0RIa3c1SkZEVWxQK2VaLwp5M2pzUUlWQ2V5Y3pXUkpuTUswK1Bwc0tUUzgvMklWL3RjK1dDUklTcHVwTU1QbVlvbnd3cXhJNHdrNWtyZkVsClM4RzhYQ1JRVmFYQ1Jud0ZJZVovWWorMFdLRHZJalljVWx6QnF4a2xBb0dBY2ZLci9QNS9wcFIrdng1S2JNd2cKSlJVeUxSeXp3NnpHelF5RnFNY2VHdnNWYXg5WnpabVRNVkpRVzNmU0xoYjd3NWtpVXpSNWloTno1T1pwcmNYUQpFUUtsQ3BTWkhIcFVWaXZoenEvM3RmRytkMmx3dHdyNElCNmVnWllxUFhpNmRWdDl4dkw1OXh6a2ErN3hlY2dVCi9lRGtleGlscWJlQ05hTEh5ZGRsWmZVQ2dZRUFzZlp1L0UyZ2pCVXJJQVFZMlFGV2pjTnJLQysyRWh0dWdTZU0KTTNoNThzeGppT0U5bVdGZnE5SEs1U3Jla282VW00cG5GZkFaOFR0bStuTytkRk5zdUd6WnpRWFhjd3J3dmdpLworVVd2a3BWYVBqTGpXUU9obkRPbkZRTlJyaDVQZ0VDbzU2NDlMWDJ5MDBwSmNlQnRlZms4eUx2cTcxYWNiekZxCmZUQ3lnZlVDZ1lCdHhWWldmU0NxUXlNcElrb2x4NlYxSHk1eExQZFBSOXQ1T2Nxa0JKdUc0b2NZUzZvR0toT3cKUDdLVWpvQld5OUJoM3hFTGxZTy9EcVZuTW5sNHJkNTJ5N2hEaVd6T0xXMWNxQmttVnl6SGhWMlhMb2xrbEdvYwpKeHpiWW0rUDBaQmM0T3NGYnZVcDZUSU03OHR5cXpna0g1TW10SE5NaUFCZFp3eFcwdnBtNlE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="

// const PUBLIC_KEY = "LS0tLS1CRUdJTiBSU0EgUFVCTElDIEtFWS0tLS0tCk1JSUJDZ0tDQVFFQXY4aTI1Q3ZhYmNnNFZ6cGNyZ2RkZDFSMVp4RStEeGdJZVpmelc2SUNSNnFOUE9zb3NXK0QKb2QvSlNlcW9FRWtBc2pPWWpEelhCZ3YzQW10V1BEYkFTSGt0UHBTVzYrWVgyN2JybnR6LzJGOVVNZGtJUStLUApmWk9HU3NCMVFpSU1MekJIOURqdGRZVDRqcTB2WVdscDZ0U1hCcUtZQXBKMFJrQVFFWU5DWjU0UGV0RUUxVnM4Cmk4dUZ6amtyeWltYzRsWEkzcmRsQXhPTEhEcHo0V0JJNSsxV0dsZ3hHNk9PR1VXaW5LNW1sckMvQ1pKTENzelcKdkREZVh4eDlRMC8xM3N6NkhKMnJ2TENpVFBZbkN1WEs5d0o2L3R5WU9DS2tUMGVFN2s0bnp6dUlsVzYxelZGVQpWMDVRQ0JuSWMzMDZZRHhiS1habnlldXJsZ21Bd0R0Zlp3SURBUUFCCi0tLS0tRU5EIFJTQSBQVUJMSUMgS0VZLS0tLS0="

// func Decrypt(data string) (string, error) {
// 	priv, err := base64.StdEncoding.DecodeString(PRIVATE_KEY)
// 	if err != nil {
// 		panic(err)
// 	}

// 	data2, err := base64.StdEncoding.DecodeString(data)
// 	if err != nil {
// 		return "", err
// 	}

// 	block, _ := pem.Decode(priv)
// 	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)

// 	if err != nil {
// 		return "", err
// 	}

// 	decrypted, err := rsa.DecryptOAEP(sha3.New512(), rand.Reader, key, data2, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(decrypted), nil
// }

// func Encrypt(msg, publicKey string) (string, error) {
// 	puclicKeyDecoded, err := base64.StdEncoding.DecodeString(publicKey)

// 	if err != nil {
// 		panic(err)
// 	}

// 	p, decoded := pem.Decode([]byte(puclicKeyDecoded))

// 	if decoded == nil {
// 		log.Fatal("Failed to decode PEM")
// 	}
// 	pub, err := x509.ParsePKCS1PublicKey(p.Bytes)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	encryptedBytes, err := rsa.EncryptOAEP(sha3.New512(), rand.Reader, pub, []byte(msg), nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
// }

func main() {
	// PAN := "01234567890"

	// encrypted, err := Encrypt(PAN, PUBLIC_KEY)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// encrypted := "KVrDBeNDvPrOQk/wNOgbQLUbJzPF2anN0rjBtvMhRyWmA84COD/oypS/u14iBWFf17tqqsuz2YCuvNu086Bp4AKgbbeqtj7HkzCS8Z9+vcmpx48ftdBvsgpES4Gtw5ajo1i7HkI01W8HWyEimM6RYFPX5VD65YDVgfLXuCvX+N4/5Kv1gJpBC1lhhJJWs8d8DEqV/618lHa79qBzZh2p2eMKx3xsP/YSO5QB5iAyCkW2VFk7nGNwTmXRjSaomnr8/I+B/ry9h8qwel3kAqWiMEjynKyJ+0fLmmKXB9QPYyy2P5AUMS1/jNZT/r9AjndWv28OQ8lMXx0akS8aIX4slw=="

	// decrypted, _ := Decrypt(encrypted)

	// println(decrypted)

	generateKeys()
}
