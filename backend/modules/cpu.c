#include <linux/module.h>
// para usar KERN_INFO
#include <linux/kernel.h>

//Header para los macros module_init y module_exit
#include <linux/init.h>
//Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
/* for copy_from_user */
#include <asm/uaccess.h>	
/* Header para usar la lib seq_file y manejar el archivo en /proc*/
#include <linux/seq_file.h>

#include <linux/sched.h>

#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Módulo de obtención de información de CPU");
MODULE_AUTHOR("Kenneth Haroldo López López");

struct task_struct* cpu;
struct task_struct* child;
struct list_head* lstProcess;

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int escribir_archivo(struct seq_file *archivo, void *v)
{   
    int ram, childram;
    char separator, childseparator;
    separator = '\0';
    childseparator = '\0';
    seq_printf(archivo, "\"procs\":[");
    for_each_process(cpu){
        seq_printf(archivo, "%c", separator);
        seq_printf(archivo, "{\"pid\":");
        seq_printf(archivo, "%d", cpu->pid);
        seq_printf(archivo, ",\"nombre\":");
        seq_printf(archivo, "%s", cpu->comm);
        seq_printf(archivo, ",\"usuario\":");
        seq_printf(archivo, "%d", cpu->real_cred->uid);
        seq_printf(archivo, ",\"estado\":");
        seq_printf(archivo, "%d", cpu->__state);
        if (cpu->mm) {
            ram = (get_mm_rss(cpu->mm)<<PAGE_SHIFT)/(1024*1024);
            seq_printf(archivo, ", \"ram\":");
            seq_printf(archivo, "%d", ram);
        }
        seq_printf(archivo, ",\"children\":[");
        childseparator = '\0'
        list_for_each(lstProcess, &(cpu->children)){
            child = list_entry(lstProcess, struct task_struct, sibling);
            seq_printf(archivo, "%c", childseparator);
            seq_printf(archivo, "{\"pid\":");
            seq_printf(archivo, "%d", child->pid);
            seq_printf(archivo, ",\"nombre\":");
            seq_printf(archivo, "%s", child->comm);
            seq_printf(archivo, ",\"usuario\":");
            seq_printf(archivo, "%d", child->real_cred->uid);
            seq_printf(archivo, ",\"estado\":");
            seq_printf(archivo, "%d", child->__state);
            if (child->mm) {
                childram = (get_mm_rss(child->mm)<<PAGE_SHIFT)/(1024*1024);
                seq_printf(archivo, ", \"ram\":");
                seq_printf(archivo, "%d", childram);
            }
            seq_printf(archivo, "}");
            childseparator = ',';
        }
        seq_printf(archivo, "]");
        separator = ',';
    }

    seq_printf(archivo, "]}");
    return 0;
}

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

//Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int _insert(void)
{
    proc_create("cpu_201906570", 0, NULL, &operaciones);
    printk(KERN_INFO "Kenneth Haroldo López López\n");
    return 0;
}

//Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("cpu_201906570", NULL);
    printk(KERN_INFO "Segundo Semestre 2022\n");
}

module_init(_insert);
module_exit(_remove);