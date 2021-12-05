package models

import (
	"net/http"

	"github.com/alen/echo-framework/db"
	"github.com/go-playground/validator/v10"
)

type Pegawai struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama" validate:"required"`
	Alamat    string `json:"alamat" validate:"required"`
	Telephone string `json:"telephone" validate:"required"`
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var ArrObj []Pegawai

	var res Response

	con := db.CreateConf()
	sqlStatment := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatment)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telephone)
		if err != nil {
			return res, err
		}
		ArrObj = append(ArrObj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = ArrObj

	return res, nil
}

func StorePegawai(nama string, alamat string, Telephone string) (Response, error) {
	var res Response

	v := validator.New()
	peg := Pegawai{
		Nama:      nama,
		Alamat:    alamat,
		Telephone: Telephone,
	}
	err := v.Struct(peg)
	if err != nil {
		return res, err
	}

	con := db.CreateConf()

	sqlStatment := "INSERT pegawai (nama, alamat, telephone ) VALUES (?, ?, ? )"
	stmt, err := con.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(nama, alamat, Telephone)
	if err != nil {
		return res, err
	}

	lastinsertedID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"Last ID": lastinsertedID,
	}
	return res, nil
}

func UpdatePegwai(id int, nama string, alamat string, telephone string) (Response, error) {
	var res Response

	con := db.CreateConf()

	sqlSatement := "UPDATE pegawai SET nama= ?, alamat = ?, telephone= ? WHERE id = ?"

	stmt, err := con.Prepare(sqlSatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telephone, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, nil
}

func DeletePegawai(id int) (Response, error) {
	var res Response

	conn := db.CreateConf()

	sqlStatment := "DELETE from pegawai WHERE id = ?"

	stmt, err := conn.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}
	rowsAfected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAfected,
	}
	return res, nil
}
