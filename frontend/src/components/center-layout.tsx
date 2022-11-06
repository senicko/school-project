export type CenterLayoutProps = {
  children: React.ReactNode;
};

export const CenterLayout = ({ children }: CenterLayoutProps) => {
  return (
    <main className="mx-auto max-w-5xl p-4 flex flex-col gap-16">
      {children}
    </main>
  );
};
