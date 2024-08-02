package writemanager

import (
	"database/sql"
	"fmt"
	"github.com/NocturnalLament/yugigo/writercommon"
)

type WriterInterface interface {
	Write() error
}

type WriterData struct {
	cardIndex                           int
	cardName                            string
	setName                             string
	cardUrl                             string
	data                                []*writercommon.GetSqlStmtInterface
	amountOfCardsToWrite                int
	structWriteType                     string
	WriteManagerStatementStringCallback func(statement *sql.Stmt)
	sqliteStatement                     []*sql.Stmt
}

func NewWriterData() *WriterData {
	return &WriterData{
		cardIndex:                           0,
		data:                                nil,
		amountOfCardsToWrite:                0,
		structWriteType:                     "",
		WriteManagerStatementStringCallback: nil,
		sqliteStatement:                     make([]*sql.Stmt, 0),
	}
}

func (w *WriterData) SetAmountOfCardsToWrite(amountOfCards int) {
	w.amountOfCardsToWrite = amountOfCards
}

func (w *WriterData) SetData(data []*writercommon.GetSqlStmtInterface) {
	w.data = data
	w.amountOfCardsToWrite = len(data)
}

func (w *WriterData) AddItemToData(item *writercommon.GetSqlStmtInterface) {
	w.data = append(w.data, item)
	w.amountOfCardsToWrite += 1
}

func (w *WriterData) SetCallback(f func(statement *sql.Stmt)) {
	w.WriteManagerStatementStringCallback = f
}

func (w *WriterData) AddSqliteStatement(statement *sql.Stmt) {
	w.sqliteStatement = append(w.sqliteStatement, statement)
}

func (w *WriterData) assignStatementToManager() {
	w.WriteManagerStatementStringCallback(w.sqliteStatement[w.cardIndex])
}

func (w *WriterData) AddSqliteStatements(statements []*sql.Stmt) {
	w.sqliteStatement = append(w.sqliteStatement, statements...)
}

type WriteManager struct {
	active               bool
	initizalized         bool
	amountOfCardsToWrite int
	data                 *WriterData
	sqliteStatement      *sql.Stmt
}

func (w *WriteManager) InitManager() {
	w.data = NewWriterData()
	w.data.WriteManagerStatementStringCallback = w.SetStatement
}

func (w *WriteManager) SetStatement(statement *sql.Stmt) {
	w.sqliteStatement = statement
}

type YGOPriceBridgeData struct {
	CardName string
	CardSet  string
}

func NewManager() *WriteManager {
	fmt.Println("Writer initialized!")
	return &WriteManager{
		active:               false,
		initizalized:         false,
		amountOfCardsToWrite: 0,
		data:                 NewWriterData(),
		sqliteStatement:      nil,
	}
}

func (w *WriteManager) SetSqliteStatement(statement *sql.Stmt) {
	w.sqliteStatement = statement
}

func (w *WriteManager) SetDataCallback() {
	w.data.WriteManagerStatementStringCallback = w.SetStatement
}
