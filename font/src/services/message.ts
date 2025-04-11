import { message } from "antd";

const [middleApi, middleContextHolder] = message.useMessage({
  top: "50%",
  transitionName: "translate(-50%, -50%)", // 居中定位
  duration: 3,
});

// 创建第二个实例（右下角位置）
const [bottomRightApi, bottomRightContextHolder] = message.useMessage({
  top: window.innerHeight - 100,
  transitionName: "translate(-50%, -50%)",
  duration: 3,
});

export interface ConfigOptions {
  top?: string | number;
  duration?: number;
  prefixCls?: string;
  getContainer?: () => HTMLElement;
  transitionName?: string;
  maxCount?: number;
  rtl?: boolean;
}

export {
  middleApi,
  middleContextHolder,
  bottomRightApi,
  bottomRightContextHolder,
};
