import React, { useCallback } from 'react';
import { useRecoilState } from 'recoil';
import Logo from '$components/logo';
import { Props as LayoutProps } from '$layouts/index';
import { screenState } from '$store/atoms/app';

type Props = Parameters<NonNullable<LayoutProps['renderAppBar']>>[0];
const Appbar: React.FC<Props> = ({ className = '', openNavigation }) => {
  const [screen] = useRecoilState(screenState);
  const { lg } = screen;

  const handleLogoClick = useCallback(
    function () {
      openNavigation();
    },
    [openNavigation]
  );

  return (
    <div className={className}>
      <div className="flex justify-center items-center h-full px-4">
        <div className="flex-none">
          {!lg && (
            <Logo
              className="h-8"
              left="text-on-complementary"
              right="text-on-complementary-variant"
              onClick={handleLogoClick}
            />
          )}
        </div>
        <div className="flex-1 min-w-0" />
      </div>
    </div>
  );
};
export default Appbar;
