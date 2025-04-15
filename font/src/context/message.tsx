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

export type setMessageNumber = (messageNumber: number | ((prev: number) => number)) => void;

const SetMessageNumberContext = createContext<setMessageNumber>(
  (messageNumber: number | ((prev: number) => number)) => {
    if (typeof messageNumber === 'function') {
      // 处理传入函数的情况，这里只是为了类型兼容，实际会报错
      messageNumber(0);
    }
    throw new Error('未在 MessageNumberContext.Provider 中使用');
  }
);

export { SetMessageNumberContext };
