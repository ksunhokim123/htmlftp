import * as React from 'react';
import { Icon, Step, Tab } from 'semantic-ui-react';

interface State {
  step: number;
}
const panes = [
  { menuItem: 'Tab 1', render: () => <Tab.Pane
  active = {false}
  ></Tab.Pane> },
  { menuItem: 'Tab 2', render: () => <Tab.Pane active={false}></Tab.Pane> },
  { menuItem: 'Tab 3', render: () => <Tab.Pane active={false}></Tab.Pane> },
];

export class RegisterStep extends React.Component<{}, State> {
  public state: any;
  constructor(props: any) {
    super(props);
    this.state = { step: 0 };
  }
  public render() {
    return (
      <Tab className='disabled' panes={panes}  attached='top'/>
    );
  }
}
