package mongo

import (
	"fmt"
	"time"

	"github.com/dlintw/goconf"
	"github.com/globalsign/mgo"
)

var (
	MongoSession *mgo.Session
)

func InitMongo(conf *goconf.ConfigFile) (err error) {

	var dialInfo = &mgo.DialInfo{}
	//GET CONF
	var addrs, _ = conf.GetString("mongo", "addrs")       //"127.0.0.1:27017"
	var timeout, _ = conf.GetInt("mongo", "timeout")      //second
	var database, _ = conf.GetString("mongo", "database") //
	var source, _ = conf.GetString("mongo", "source")
	var username, _ = conf.GetString("mongo", "username")
	var password, _ = conf.GetString("mongo", "password")
	var poollimit, _ = conf.GetInt("mongo", "poollimit") //Session.SetPoolLimit
	//Dail DB
	dialInfo.Addrs = append([]string{}, addrs)
	dialInfo.Direct = false
	dialInfo.Timeout = time.Second * time.Duration(timeout)
	dialInfo.Database = database
	dialInfo.Source = source
	dialInfo.Username = username
	dialInfo.Password = password
	dialInfo.PoolLimit = poollimit
	//fmt.Printf("---%v--", dialInfo)
	MongoSession, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	//defer MongoSession.Close()  move to invoke part
	//Switch the session to a monotonic behavior.
	MongoSession.SetMode(mgo.Monotonic, true)
	return
}
