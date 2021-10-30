import Alert from "react-bootstrap/Alert";

function alert() {
  return (
    <Alert
      style={{
        marginTop: "20%",
        textAlign: "center",
        color: "white",
        backgroundColor: "cadetblue",
      }}
    >
      <b>No data found.</b>
    </Alert>
  );
}
export default alert;
