"use client";

import { Badge } from "@/components/ui/badge";
import {
  Table,
  TableHeader,
  TableBody,
  TableHead,
  TableRow,
  TableCell,
} from "@/components/ui/table";
import { useEffect, useState } from "react";
import axios from "axios";
import { LogType } from "@/config/Types";

const LogsPage = () => {
  //   const lgs = [
  //     {
  //       clientIp: "192.168.1.139",
  //       domain: "google.com",
  //       port: "12345",
  //       type: "HTTPS",
  //       timeStamp: new Date().toISOString(),
  //     },
  //     {
  //       clientIp: "192.168.1.139",
  //       domain: "google.com",
  //       port: "12345",
  //       type: "HTTPS",
  //       timeStamp: new Date().toISOString(),
  //     },
  //     {
  //       clientIp: "192.168.1.139",
  //       domain: "google.com",
  //       port: "12345",
  //       type: "HTTPS",
  //       timeStamp: new Date().toISOString(),
  //     },
  //     {
  //       clientIp: "192.168.1.139",
  //       domain: "google.com",
  //       port: "12345",
  //       type: "HTTPS",
  //       timeStamp: new Date().toISOString(),
  //     },
  //   ];

  const [logs, SetLogs] = useState<LogType[]>([]);
  useEffect(() => {
    (async () => {
      const lgs = await axios.get(`http://localhost:4004/api/logs`);
      SetLogs(lgs.data);
    })();
  }, []);

  return (
    <>
      <div className="w-full h-screen overflow-y-scroll no-scrollbar p-6 overflow-auto">
        <Table className="w-full text-sm">
          <TableHeader>
            <TableRow className="border-b">
              <TableHead className="text-left py-2">ID</TableHead>
              <TableHead className="text-left py-2">Client</TableHead>
              <TableHead className="text-left py-2">Domain</TableHead>
              <TableHead className="text-left py-2">Type</TableHead>
              <TableHead className="text-left py-2">Timestamp</TableHead>
            </TableRow>
          </TableHeader>

          <TableBody>
            {logs.map((log) => (
              <TableRow key={log.id}>
                <TableCell className="py-2">{log.id}</TableCell>

                <TableCell className="py-2">
                  <Badge variant="secondary" className="rounded-sm">
                    {log.client_ip}
                  </Badge>
                </TableCell>

                <TableCell className="py-2 font-medium">{log.domain}</TableCell>

                <TableCell className="py-2">
                  <Badge
                    variant={
                      log.query_type === "A"
                        ? "default"
                        : log.query_type === "AAAA"
                          ? "secondary"
                          : "outline"
                    }
                    className="rounded-sm"
                  >
                    {log.query_type}
                  </Badge>
                </TableCell>

                <TableCell className="py-2 text-muted-foreground">
                  {new Date(log.timestamp).toLocaleString()}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </>
  );
};

export default LogsPage;
