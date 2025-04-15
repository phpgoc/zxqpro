import UserInProject from "../../components/userInProject.tsx";
import { useParams } from "react-router-dom";
import { parseIdToNumber } from "../../services/utils.ts";

export default function ProjectView() {
  const { id } = useParams();
  let projectNumericId: number = parseIdToNumber(id);
  return (
    <>
      <UserInProject id={projectNumericId} />
    </>
  )
}
