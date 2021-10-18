import classnames from 'classnames';
import React from 'react';
import { On } from '$constants/index';
import { ClassName } from '$types/index';

export type Props = {
  on: On;
  className?: ClassName;
};
const Spinner: React.FC<Props> = ({ on, className = '' }) => {
  const ballClassNames = classnames(
    'absolute w-[40%] h-[40%] transform animate-pulse'
  );
  const ballInnerClassNames = classnames(
    'w-full h-full rounded-full animate-bounce'
  );
  return (
    <div className={classnames(className)}>
      <div className="aspect-w-1 aspect-h-1">
        <div className="animate-spin-slow">
          <div
            className={classnames(
              'top-0 left-0 rotate-[135deg]',
              ballClassNames
            )}
          >
            <div
              className={classnames(ballInnerClassNames, `bg-on-${on}-high`)}
            />
          </div>
          <div
            className={classnames(
              'top-0 right-0 rotate-[225deg]',
              ballClassNames
            )}
          >
            <div
              className={classnames(ballInnerClassNames, `bg-on-${on}-medium`)}
            />
          </div>
          <div
            className={classnames(
              'bottom-0 right-0 rotate-[315deg]',
              ballClassNames
            )}
          >
            <div
              className={classnames(ballInnerClassNames, `bg-on-${on}-low`)}
            />
          </div>
          <div
            className={classnames(
              'bottom-0 left-0 rotate-[45deg]',
              ballClassNames
            )}
          >
            <div
              className={classnames(ballInnerClassNames, `bg-on-${on}-faint`)}
            />
          </div>
        </div>
      </div>
    </div>
  );
};
export default Spinner;
