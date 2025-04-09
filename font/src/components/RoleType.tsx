import {Select} from "antd";
import {roleTypesMap} from "../types/project.ts";

export default function RoleTypeSelect({ roleType, onChange }: { roleType: string; onChange: (newStatus: string) => void }) {


  return (
      <Select
          placeholder="Select role type"
          onChange={onChange}
          style={{ width: 200 }}
      >
          <Select.Option key={0} value={0}>
                All
          </Select.Option>
            {Object.entries(roleTypesMap).map(([key, value]) => (
                <Select.Option key={key} value={key} selected={roleType === key}>
                    {value}
                </Select.Option>
            ))}

      </Select>
  );
}