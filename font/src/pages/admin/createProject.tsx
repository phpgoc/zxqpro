import UserListSelect from "../../components/UserList.tsx";
import { useState } from "react";
export default function CreateProject() {
  const [userId, setUserId] = useState(0);
  return (
    <div>
      <UserListSelect userId={userId} onChange={setUserId} />
      hello world
    </div>
  );
}
