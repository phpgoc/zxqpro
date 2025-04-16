import MarkdownViewer from "../../components/markdown.tsx";
import { Select, Space } from "antd";
import { useState } from "react";

export default function Help() {
  const items = {
    ["role"]: "角色解释",
    ["develop_env"]: "开发环境",
    ["develop_command"]: "开发命令",
  };
  const [markdownKey, setMarkdownKey] = useState("role");

  return (<>
      <Space direction="horizontal" style={{ marginBottom: 16 }}>
        <Select
          placeholder="Select markdown"
          onChange={(v) => setMarkdownKey(v)}
          style={{ width: 200 }}
        >
          {
            Object.entries(items).map(
              ([key, value]) => {
                return (
                  <Select.Option key={key} value={key} selected={key === markdownKey}>
                    {value}
                  </Select.Option>
                );
              }
            )
          }
        </Select>
      </Space>
    {/*<div style={{*/}
    {/*  backgroundColor: '#2d2d2d', // 设置深色背景*/}
    {/*  color: '#ffffff', // 设置文本颜色为白色，确保与背景有足够对比度*/}
    {/*  padding: '20px', // 添加内边距，使内容不紧贴边缘*/}
    {/*  borderRadius: '8px' // 添加圆角，让视觉效果更柔和*/}
    {/*}}>*/}
      <MarkdownViewer filePath={`${markdownKey}.md`} />
    {/*</div>*/}
    </>
  );
}