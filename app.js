const { exec } = require("child_process");
const path = require("path");

const goExePath = path.join(__dirname, "main.exe");

exec(goExePath, (error, stdout, stderr) => {
  if (error) {
    console.error(`Error executing the Go executable: ${error.message}`);
    return;
  }

  if (stderr) {
    console.error(`Go executable returned an error: ${stderr}`);
    return;
  }

  const goOutput = stdout.trim();
  console.log(`Output of the Go executable: ${goOutput}`);
});
