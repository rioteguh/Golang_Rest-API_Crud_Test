package models

import (
	"database/sql"
	"entities"
)

type NamaModel struct{
	Db *sql.DB
}

func (namaModel NamaModel) FindAll() (nama []entities.Nama, err error){
	rows, err := namaModel.Db.Query("select * from tbl_test")
	if err != nil {
		return nil, err
	} else {
		var namas []entities.Nama
		for rows.Next(){
			var id int64
			var name string
			err2 := rows.Scan(&id, &name)
			if err2 != nil {
				return nil, err2
			} else {
				nama := entities.Nama{
					Id: id,
					Name: name,
				}
				namas = append(namas, nama)
			}
		}
		return namas, nil
	}
}

func (namaModel NamaModel) Search(keyword string) (nama []entities.Nama, err error){
	rows, err := namaModel.Db.Query("select * from tbl_test where name like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		var namas []entities.Nama
		for rows.Next(){
			var id int64
			var name string
			err2 := rows.Scan(&id, &name)
			if err2 != nil {
				return nil, err2
			} else {
				nama := entities.Nama{
					Id: id,
					Name: name,
				}
				namas = append(namas, nama)
			}
		}
		return namas, nil
	}
}

func (namaModel NamaModel) Create(nama *entities.Nama) (err error){
	result, err := namaModel.Db.Exec("insert into tbl_test(name) values (?)", nama.Name)
	if err != nil {
		return err
	} else {
		nama.Id, _ = result.LastInsertId()
		return nil
	}
}

func (namaModel NamaModel) Update(nama *entities.Nama) (int64, error){
	result, err := namaModel.Db.Exec("update tbl_test set name = ? where id = ?", nama.Name, nama.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (namaModel NamaModel) Delete(id int64) (int64, error){
	result, err := namaModel.Db.Exec("delete from tbl_test where id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}