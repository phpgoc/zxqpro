import MarkdownViewer from "../../components/markdown.tsx";
import { Select, Space } from "antd";
import { useState } from "react";

export default function Help() {
  const items = {
    ["role"]: "角色",
    ["develop"]: "开发"
  };
  const [markdownKey, setMarkdownKey] = useState("role");

  return (<>
      <Space direction="horizontal" style={{ marginBottom: 16 }}>
        <Select
          placeholder="Select role type"
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
      <MarkdownViewer filePath={`${markdownKey}.md`} />
    </>
  );
}