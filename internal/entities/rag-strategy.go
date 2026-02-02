package entities

type RAGStrategy string

const (
	RAGWebhook RAGStrategy = "webhook"
	RAGCron    RAGStrategy = "cron"
	RAGManual  RAGStrategy = "manual"
)
