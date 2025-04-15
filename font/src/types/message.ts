export  type SSEMessage = {
  code: number
  message: string
  link?: string
}

export type Message = {
  id:     number
  user_name: string
  message : string
  link ?:    string
  time    : string
  read     :boolean
}