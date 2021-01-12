import React , {useEffect} from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import YouTube from '@material-ui/icons/YouTube';
import SignOut from '@material-ui/icons/Settings';


import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';



export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    
    <SidebarItem
            icon={CreateComponentIcon}
      to="Doctorinfo"
      text="ข้อมูลแพทย์"
    />
    <SidebarItem
      icon={CreateComponentIcon}
      to="createPatientrecord"
      text="ลงทะเบียนผู้ป่วย"
    />
    <SidebarItem
      icon={CreateComponentIcon}
      to="create_Patientrights"
      text="สิทธิผู้ป่วย"
    />
    <SidebarItem
            icon={CreateComponentIcon}
      to="createHistorytaking"
      text="ซักประวัติ"
    />
    <SidebarItem
            icon={CreateComponentIcon}
      to="Treatment"
      text="บันทึกการรักษา"
    />
    <SidebarItem
            icon={CreateComponentIcon}
      to="createbill"
      text="ส่วนการเงิน"
    />

    <SidebarPinButton />

  <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      to="sign_out"
      text="Sign Out"
    />

   
  </Sidebar>
);
