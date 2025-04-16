import UserInProjectSelect from "../../components/userInProjectSelect.tsx";
import { useParams } from "react-router-dom";
import { parseIdToNumber } from "../../services/utils.ts";

export default function ProjectView() {
  const { id } = useParams();
  let projectNumericId: number = parseIdToNumber(id);
  return (
    <>
      <UserInProjectSelect id={projectNumericId} />
    </>
  )
}
