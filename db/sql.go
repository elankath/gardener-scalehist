package db

const CreateWorkerPoolInfo = `CREATE TABLE IF NOT EXISTS worker_pool_info(
	RowID INTEGER PRIMARY KEY AUTOINCREMENT,
	CreationTimestamp INT NOT NULL,
	SnapshotTimestamp INT NOT NULL,
	Name TEXT,
	Namespace TEXT,
	MachineType      TEXT, 
	Architecture      TEXT,
	Minimum           INT,
	Maximum           INT,
	MaxSurge          TEXT,
	MaxUnavailable    TEXT,
	Zones             TEXT,
	DeletionTimestamp DATETIME,
	Hash TEXT)`

const InsertWorkerPoolInfo = `INSERT INTO worker_pool_info(
	CreationTimestamp,
	SnapshotTimestamp,
	Name,
	Namespace,
	MachineType,
	Architecture,
	Minimum,                       
	Maximum,
	MaxSurge,
	MaxUnavailable ,
	Zones,
	Hash
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

const SelectWorkerPoolInfoBefore = `SELECT * from worker_pool_info where SnapshotTimestamp <= ? ORDER BY SnapshotTimestamp DESC`

const CreateMCDInfoTable = `CREATE TABLE IF NOT EXISTS mcd_info(
	RowID INTEGER PRIMARY KEY AUTOINCREMENT,
	CreationTimestamp INT NOT NULL,
	SnapshotTimestamp INT NOT NULL,
	Name TEXT,
	Namespace TEXT,
	Replicas INTEGER,
	PoolName TEXT,
	Zone TEXT,
	MaxSurge TEXT,
	MaxUnavailable TEXT, 
	MachineClassName TEXT,
	DeletionTimestamp DATETIME,
	Hash TEXT)`
const InsertMCDInfo = `INSERT INTO mcd_info(
	CreationTimestamp,
	SnapshotTimestamp,
	Name,
	Namespace,
	Replicas,
	PoolName,
	Zone,
	MaxSurge,
	MaxUnavailable,
	MachineClassName,
	Hash) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
const SelectMCDBefore = `SELECT * from mcd_info where SnapshotTimestamp <= ? ORDER BY SnapshotTimestamp DESC`
const UpdateMCDInfoDeletionTimestamp = `UPDATE mcd_info SET DeletionTimestamp = ? where Name = ?`
const SelectMCDInfoHash = "SELECT Hash FROM mcd_info WHERE name=? ORDER BY RowID desc LIMIT 1"
const SelectLatestMCDInfo = "SELECT * FROM mcd_info WHERE name=? ORDER BY RowID DESC LIMIT 1"

const CreateNodeInfoTable = `CREATE TABLE IF NOT EXISTS node_info (
	RowID INTEGER PRIMARY KEY AUTOINCREMENT,
	CreationTimestamp INT NOT NULL,
	SnapshotTimestamp INT NOT NULL,
	Name TEXT, 
	Namespace TEXT, 
	ProviderID TEXT, 
	AllocatableVolumes INTEGER,
	Labels TEXT, 
	Taints TEXT, 
	Allocatable TEXT, 
	Capacity TEXT, 
	DeletionTimestamp DATETIME,
	Hash TEXT)`

const InsertNodeInfo = `INSERT INTO node_info(
    CreationTimestamp,
	SnapshotTimestamp,
	Name,
	Namespace,
	ProviderID,
    AllocatableVolumes,
	Labels,
	Taints,
	Allocatable,
	Capacity,
	Hash) 
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

const SelectNodeInfoBefore = `SELECT * FROM node_info WHERE CreationTimestamp < ? ORDER BY CreationTimestamp DESC`
const SelectNodeCountWithNameAndHash = "SELECT COUNT(*) from node_info where Name=? and Hash=?"
const UpdateNodeInfoDeletionTimestamp = `UPDATE node_info SET DeletionTimestamp = ? where Name = ?`

const CreatePodInfoTable = `CREATE TABLE IF NOT EXISTS pod_info (
	RowID INTEGER PRIMARY KEY AUTOINCREMENT,
	CreationTimestamp INT NOT NULL,
	SnapshotTimestamp INT NOT NULL,
	Name TEXT,
	Namespace TEXT,
	UID TEXT NOT NULL,
	NodeName TEXT,
	NominatedNodeName TEXT,
	Labels TEXT,
	Requests TEXT,
	Spec TEXT,
	ScheduleStatus INTEGER,
	DeletionTimestamp INT,
	Hash TEXT)`
const InsertPodInfo = `INSERT INTO pod_info(
    CreationTimestamp,
	SnapshotTimestamp,
	Name, 
	Namespace,
	UID, 
	NodeName,
    NominatedNodeName,
	Labels,
	Requests,
    Spec,
    ScheduleStatus,
	Hash) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
const UpdatePodDeletionTimestamp = "UPDATE pod_info SET DeletionTimestamp=? WHERE UID=?"
const SelectPodCountWithUIDAndHash = "SELECT COUNT(*) from pod_info where UID=? and Hash=?"
const SelectPodsWithEmptyNameAndBeforeCreationTimestamp = `SELECT * FROM (SELECT * from pod_info
               WHERE ScheduleStatus = 0 AND CreationTimestamp <= ? AND (DeletionTimestamp is null OR DeletionTimestamp >= ?)
               ORDER BY CreationTimestamp DESC) GROUP BY Name;`

const SelectLatestScheduledPodsBeforeSnapshotTimestamp = `SELECT * from (SELECT * FROM pod_info WHERE (ScheduleStatus = 1)  
                AND SnapshotTimestamp <= ? AND (DeletionTimestamp is null OR DeletionTimestamp >=  ?)  ORDER BY SnapshotTimestamp DESC) 
                GROUP BY Name;`
const SelectLatestPodsBeforeSnapshotTimestamp = `SELECT * FROM pod_info WHERE
                SnapshotTimestamp <= ? AND (DeletionTimestamp is null OR DeletionTimestamp >=  ?)  ORDER BY SnapshotTimestamp DESC;`

const CreateEventInfoTable = `CREATE TABLE IF NOT EXISTS event_info(
	UID varchar(128) PRIMARY KEY,
	EventTime DATETIME NOT NULL,
	ReportingController VARCHAR(256),
	Reason VARCHAR(128),
	Message TEXT,
	InvolvedObjectKind varchar(128),
	InvolvedObjectName varchar(128),
	InvolvedObjectNamespace varchar(128),
	InvolvedObjectUID varchar(128))`

const InsertEvent = `INSERT INTO event_info(
	UID,
	EventTime,
	ReportingController,
	Reason,
	Message,
	InvolvedObjectKind,
	InvolvedObjectName,
	InvolvedObjectNamespace,
	InvolvedObjectUID) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?) ON CONFLICT(UID) DO NOTHING`

const CreateCASettingsInfoTable = `CREATE TABLE IF NOT EXISTS ca_settings_info(
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Expander TEXT,
    MaxNodesTotal INT,
	Priorities TEXT,
	Hash TEXT)`

const SelectCADeploymentByHash = `SELECT * FROM ca_settings_info WHERE Hash=?`

const SelectLatestCADeployment = `SELECT * FROM ca_settings_info ORDER BY Id DESC LIMIT 1`

const InsertCADeployment = `INSERT INTO ca_settings_info (
    Expander,
    MaxNodesTotal,
	Priorities,
    Hash
) VALUES (? ,? , ? ,?)`

const SelectLatestNodesBeforeAndNotDeleted = `SELECT * from (select * from node_info where node_info.CreationTimestamp <= ? and node_info.DeletionTimestamp = 0 order by RowID DESC) GROUP by Name`
