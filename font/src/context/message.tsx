import { createContext } from "react";
import type { MessageInstance } from "antd/es/message/interface";

// 定义上下文类型：包含所有需要全局共享的消息 API
export interface MessageContextValue {
  middleMessageApi: MessageInstance; // 居中消息 API
  bottomRightMessageApi: MessageInstance; // 右下角消息 API
}

// 创建消息上下文
const MessageContext = createContext<MessageContextValue | null>({
  middleMessageApi: {} as MessageInstance,
  bottomRightMessageApi: {} as MessageInstance,
});
export default MessageContext;
