import { createContext } from 'react';
import type { MessageInstance } from 'antd/es/message/interface';

// 定义上下文类型：包含所有需要全局共享的消息 API
interface MessageContextValue {
    middleApi: MessageInstance; // 居中消息 API
    bottomRightApi: MessageInstance; // 右下角消息 API
}

// 创建消息上下文
const MessageContext = createContext<MessageContextValue | null>(
    {
        middleApi: {} as MessageInstance,
        bottomRightApi: {} as MessageInstance,
    }
);
export default MessageContext;

