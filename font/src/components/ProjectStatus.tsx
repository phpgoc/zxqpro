import {Select} from "antd";
import {projectStatusMap} from "../types/project.ts";

export default function ProjectStatusSelect({ status, onChange }: { status: string; onChange: (newStatus: string) => void }) {


  return (
      <Select
          placeholder="Select status"
          onChange={onChange}
          style={{ width: 200 }}
      >
          <Select.Option key={0} value={0}>
                All
          </Select.Option>
            {Object.entries(projectStatusMap).map(([key, value]) => (
                <Select.Option key={key} value={key} selected={status === key}>
                    {value}
                </Select.Option>
            ))}

      </Select>
  );
}