package starter

import (
	"context"
	"sync"

	"github.com/fladago/gbbackone/app/repos/user"
)

type App struct {
	us *user.Users
}

func NewApp(ust user.UserStore) *App {
	a := &App{
		us: user.NewUsers(nil),
	}
	return a
}

//Нужно создать все объекты бизнес логики. Как минимум, все объекты, связанные с бизнес логикой
//Необходимо добавить api, которое будет посылать запросы
func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

}
