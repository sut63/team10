import React ,{useState,useEffect}from 'react';
import Backdrop from '@material-ui/core/Backdrop';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntPatientrecord } from '../../api/models';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {Grid, Modal,Typography} from '@material-ui/core';
import Fade from '@material-ui/core/Fade';
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    modal: {
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    },
    paper: {
      backgroundColor: theme.palette.background.paper,
      border: '2px solid #000',
      boxShadow: theme.shadows[5],
      padding: theme.spacing(2, 4, 3),
    },
  }),
);
export default function MoreInfo( id : any) {
  const classes = useStyles();
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };
  const http = new DefaultApi();
  const [loading, setLoading] = React.useState(true);
  const [Patientrecord, setPatientrecord] = useState<EntPatientrecord[]>(Array);
  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({offset : 0})
    setLoading(false);
    setPatientrecord(res);
  };
  useEffect(() => {
    getPatientrecord();
  }, [loading]);

  return (
    <div>
      <Button variant='outlined' color = 'primary' onClick={handleOpen}>
        รายละเอียดเพิ่มเติม
      </Button>
      <Modal
        aria-labelledby="transition-modal-title"
        aria-describedby="transition-modal-description"
        className={classes.modal}
        open={open}
        onClose={handleClose}
        closeAfterTransition
        BackdropComponent={Backdrop}
      >
        <Fade in={open}>
          <div className={classes.paper}>
              {Patientrecord.filter(p => p.id === id.id).map(patientrecord => ( 
                    <Grid>
                    <Typography variant = 'h4'>
                        ข้อมูลผู้ป่วย<br/>
                    </Typography>
                    <Typography variant = 'subtitle1'> 
                        รหัสผู้ป่วย : &emsp;{patientrecord.id}<br/>
                        คำนำหน้าชื่อ : &emsp;{patientrecord.edges?.edgesOfPrename?.prefix}<br/>
                        ชื่อ-นามสกุล :&emsp;{patientrecord.name}<br/>
                        เพศ : &emsp;{patientrecord.edges?.edgesOfGender?.genderstatus}<br/>
                        เลขบัตรประจำตัวประชาชน : &emsp;{patientrecord.idcardnumber}<br/>
                        อายุ : &emsp;{patientrecord.age}<br/>
                        กรุ๊ปเลือด : &emsp;{patientrecord.edges?.edgesOfBloodtype?.bloodtype}<br/>
                        โรคประจำตัว : &emsp;{patientrecord.disease}<br/>
                        ยาที่แพ้ : &emsp;{patientrecord.allergic}<br/>
                        เบอร์โทรที่ติดต่อได้ : &emsp;{patientrecord.phonenumber}<br/>
                        อีเมล์ : &emsp;{patientrecord.email}<br/>
                        ที่อยู่ : &emsp;{patientrecord.home}<br/>
                        วันเวลาที่ลงทะเบียน : &emsp;{patientrecord.date}<br/>
                        พนักงานเวชระเบียนที่ลงบันทึก : &emsp;{patientrecord.edges?.edgesOfMedicalrecordstaff?.name}<br/>
                    </Typography>
                    </Grid>
                       ))}
          </div>
        </Fade>
      </Modal>
    </div>
  );
}