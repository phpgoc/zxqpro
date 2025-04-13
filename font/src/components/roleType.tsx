import { Select } from "antd";
import { RoleType, roleTypesMap } from "../types/project.ts";

export default function RoleTypeSelect({
                                         roleType,
                                         onChange
                                       }: {
  roleType: number;
  onChange: (newStatus: number) => void;
}) {
  return (
    <Select
      placeholder="Select role type"
      onChange={onChange}
      style={{ width: 200 }}
    >
      <Select.Option key={0} value={0}>
        All
      </Select.Option>
      {Object.entries(roleTypesMap)
        .map(([key, value]) => (
          <Select.Option key={key} value={key} selected={roleType === Number(key)}>
            {value}
          </Select.Option>
        ))}
    </Select>
  );
}

export function RoleTypeForAddSelect({
                                       roleType,
                                       onChange
                                     }: {
  roleType: number;
  onChange: (newStatus: number) => void;
}) {
  return (
    <Select
      placeholder="Select role type"
      onChange={(v) => onChange(Number(v))}
      style={{ width: 200 }}
    >

      {Object.entries(roleTypesMap).filter(([key, _value]) => {
        const numericKey = Number(key);
        return numericKey != RoleType.ADMIN && numericKey != RoleType.OWNER;
      })
        .map(([key, value]) => (
          <Select.Option key={key} value={key} selected={roleType === Number(key)}>
            {value}
          </Select.Option>
        ))}
    </Select>
  );
}

