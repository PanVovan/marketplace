package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type OrderController struct {
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (oc *OrderController) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Get All"))
}

func (oc *OrderController) AdminGetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Admin Get one " + id))
}

func (oc *OrderController) AdminGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Admin Get user " + id))
}

func (oc *OrderController) AdminCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Admin Create "))
}

func (oc *OrderController) AdminDelete(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Admin Delete "))
}

func (oc *OrderController) UserGetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Get All"))
}

func (oc *OrderController) UserGetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("User Get one " + id))
}

func (oc *OrderController) UserCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("User Create "))
}

func (oc *OrderController) GuestCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Guest Create "))
}

/*
 * только для администратора магазина
 */

// // получить список всех заказов магазина
// router.get(
//     '/admin/getall',
//     authMiddleware, adminMiddleware,
//     OrderController.adminGetAll
// )
// // получить список заказов пользователя
// router.get(
//     '/admin/getall/user/:id([0-9]+)',
//     authMiddleware, adminMiddleware,
//     OrderController.adminGetUser
// )
// // получить заказ по id
// router.get(
//     '/admin/getone/:id([0-9]+)',
//     authMiddleware, adminMiddleware,
//     OrderController.adminGetOne
// )
// // создать новый заказ
// router.post(
//     '/admin/create',
//     authMiddleware, adminMiddleware,
//     OrderController.adminCreate
// )
// // удалить заказ по id
// router.delete(
//     '/admin/delete/:id([0-9]+)',
//     authMiddleware, adminMiddleware,
//     OrderController.adminDelete
// )
