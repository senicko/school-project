import { LearningSetsList } from "../components/learning-set/learning-sets-list";
import { Header } from "../components/header";
import { CenterLayout } from "../components/center-layout";

export const Root = () => {
  return (
    <CenterLayout>
      <Header />
      <LearningSetsList />
    </CenterLayout>
  );
};
