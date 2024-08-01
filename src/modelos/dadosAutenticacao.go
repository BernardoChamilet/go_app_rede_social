package modelos

//DadosAutenticacao contém o id e token do usuário logado/autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
